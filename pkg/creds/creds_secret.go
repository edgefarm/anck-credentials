package creds

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	namespace         = "edgefarm-network"
	credentialsSecret = "ngsaccounts"
	fixedUsername     = "customer"
)

type CredsSecrets struct {
	Creds
	client *kubernetes.Clientset
}

func NewCredsSecrets() *CredsSecrets {
	c, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(c)
	if err != nil {
		panic(err.Error())
	}
	return &CredsSecrets{
		Creds: Creds{
			Credentials: make(map[string]map[string]string),
		},
		client: clientset,
	}
}

func (c *CredsSecrets) DesiredState(account string, usernames []string) (map[string]string, error) {
	secrets, err := c.client.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("natsfunction=users, natsaccount=%s", account),
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
	return res, nil
}

func (c *CredsSecrets) DeleteAccount(account string) error {
	return nil
}
