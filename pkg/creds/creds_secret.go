package creds

import (
	"context"
	"encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	namespace         = "edgefarm-network"
	stateSecret       = "credsmanagerstate"
	credentialsSecret = "ngsaccounts"
	fixedUsername     = "customer"
)

// State stores the current state of the credentials (all used accounts)
type State struct {
	// UsedAccounts slice of used accounts
	UsedAccounts []string `json:"usedAccounts"`
}

// CredsSecrets implements CredsIf using Kubernetes secrets
type CredsSecrets struct {
	Creds
	client *kubernetes.Clientset
	State  *State
}

// NewCredsSecrets creates a new CredsSecrets
func NewCredsSecrets() *CredsSecrets {
	c, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(c)
	if err != nil {
		panic(err.Error())
	}
	var state *State
	state, err = ReadState(clientset)
	if err != nil {
		err := CreateState(clientset)
		if err != nil {
			panic(err.Error())
		}
		state, err = ReadState(clientset)
		if err != nil {
			panic(err.Error())
		}
	}

	return &CredsSecrets{
		Creds: Creds{
			Credentials: make(map[string]map[string]string),
		},
		client: clientset,
		State:  state,
	}
}

// ReadState reads the state of the credentials and returns struct State
func ReadState(client *kubernetes.Clientset) (*State, error) {
	secret, err := client.CoreV1().Secrets(namespace).Get(context.TODO(), stateSecret, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("ReadState: cannot access state secret, err: %v", err)
	}
	state := &State{}
	err = json.Unmarshal(secret.Data["state"], &state)
	if err != nil {
		return nil, fmt.Errorf("ReadState: cannot unmarshal state, err: %v", err)
	}

	return state, nil
}

// CreateState creates an empty state of the credentials
func CreateState(client *kubernetes.Clientset) error {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      stateSecret,
			Namespace: namespace,
		},
		Data: map[string][]byte{
			"state": []byte("{}"),
		},
	}
	_, err := client.CoreV1().Secrets(namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("cannot create state secret, err: %v", err)
	}
	return nil
}

// UpdateState updates the state of the credentials
func UpdateState(client *kubernetes.Clientset, state *State) error {
	secret, err := client.CoreV1().Secrets(namespace).Get(context.TODO(), stateSecret, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("UpdateState: cannot access state secret, err: %v", err)
	}
	j, err := json.Marshal(*state)
	if err != nil {
		return fmt.Errorf("UpdateState: cannot marshal state, err: %v", err)
	}
	secret.Data["state"] = j
	_, err = client.CoreV1().Secrets(namespace).Update(context.TODO(), secret, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("cannot update secret %s, err: %v", stateSecret, err)
	}
	return nil
}

// DesiredState constructs the desired state of the credentials
func (c *CredsSecrets) DesiredState(account string, usernames []string) (map[string]string, error) {
	secrets, err := c.client.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("natsaccount=%s,natsfunction=users", account),
	})

	if err != nil {
		return nil, fmt.Errorf("cannot access secrets")
	}

	if len(secrets.Items) == 0 {
		return nil, fmt.Errorf("no secrets found for account %s", account)
	}

	res := map[string]string{}
	for _, user := range usernames {
		res[user] = string(secrets.Items[0].Data[fixedUsername])
	}

	// reads the current state and checks if the account is already used.
	// if not, it is added to the used accounts slice
	// if already used return with error
	state, err := ReadState(c.client)
	if err != nil {
		return nil, fmt.Errorf("cannot read state, err: %v", err)
	}
	for _, usedAccount := range state.UsedAccounts {
		if usedAccount == account {
			return nil, fmt.Errorf("account %s already used", account)
		}
	}
	state.UsedAccounts = append(state.UsedAccounts, account)

	// update the state
	err = UpdateState(c.client, state)
	if err != nil {
		return nil, fmt.Errorf("cannot update state, err: %v", err)
	}

	return res, nil
}

// DeleteAccount deletes the account from the credentials. Not needed in this implementation.
func (c *CredsSecrets) DeleteAccount(account string) error {
	// check if the account is used.
	// if used delete it
	state, err := ReadState(c.client)
	if err != nil {
		return fmt.Errorf("cannot read state, err: %v", err)
	}
	for i, usedAccount := range state.UsedAccounts {
		if usedAccount == account {
			state.UsedAccounts = append(state.UsedAccounts[:i], state.UsedAccounts[i+1:]...)
			break
		}
	}
	err = UpdateState(c.client, state)
	if err != nil {
		return fmt.Errorf("cannot update state, err: %v", err)
	}

	return nil
}
