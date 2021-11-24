/*
Copyright © 2021 Ci4Rail GmbH <engineering@ci4rail.com>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package secrets

import (
	"context"
	"fmt"

	"github.com/edgefarm/edgefarm.network/pkg/creds"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	namespace     = "edgefarm-network"
	stateSecret   = "credsmanagerstate"
	fixedUsername = "user"
)

// CredsSecrets implements CredsIf using Kubernetes secrets
type CredsSecrets struct {
	creds.Creds
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
		Creds: creds.Creds{
			Credentials: make(map[string]map[string]string),
		},
		client: clientset,
		State:  state,
	}
}

func (c *CredsSecrets) getUnusedNatsAccounts(state *State) ([]string, error) {
	unusedNatsAccounts := []string{}
	secrets, err := c.client.CoreV1().Secrets(namespace).List(context.Background(), metav1.ListOptions{
		LabelSelector: "natsfunction=users",
	})
	if err != nil {
		return nil, fmt.Errorf("getUnusedNatsAccounts: cannot list secrets, err: %v", err)
	}

	for _, secret := range secrets.Items {
		unusedNatsAccounts = append(unusedNatsAccounts, secret.Name)
	}

	// Remove account from unusedNatsAccounts if already present in state.UsedAccounts
	for _, secret := range secrets.Items {
		for _, s := range state.UsedAccounts {
			if s.Account == secret.Name {
				unusedNatsAccounts, err = removeAccount(unusedNatsAccounts, s.Account)
				if err != nil {
					return nil, fmt.Errorf("getUnusedNatsAccounts: cannot remove account %s from unusedNatsAccounts, err: %v", s.Account, err)
				}
			}
		}
	}
	return unusedNatsAccounts, nil
}

// AllocateNatsAccount allocates a new nats account if possible
func (c *CredsSecrets) AllocateNatsAccount(applicationName string) (string, error) {
	state, err := ReadState(c.client)
	if err != nil {
		return "", fmt.Errorf("AllocateNatsAccount: cannot read state, err: %v", err)
	}
	unusedNatsAccounts, err := c.getUnusedNatsAccounts(state)
	if err != nil {
		return "", fmt.Errorf("AllocateNatsAccount: cannot get unused nats accounts, err: %v", err)
	}

	// Use the first unused account
	if len(unusedNatsAccounts) > 0 {
		account := unusedNatsAccounts[0]
		state.UsedAccounts = append(state.UsedAccounts, NatsAccountMapping{
			Account:         account,
			ApplicationName: applicationName,
		})
		err := UpdateState(c.client, state)
		if err != nil {
			return "", err
		}
		return account, nil
	} else {
		return "", fmt.Errorf("AllocateNatsAccount: no nats account available")
	}

}

func removeAccount(slice []string, s string) ([]string, error) {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...), nil
		}
	}
	return nil, fmt.Errorf("removeAccount: account %s not found", s)
}

// DesiredState constructs the desired state of credentials for a given application name
func (c *CredsSecrets) DesiredState(applicationName string, usernames []string) (map[string]string, error) {
	var natsAccount = ""

	// Check current state if application name is already used
	state, err := ReadState(c.client)
	if err != nil {
		return nil, fmt.Errorf("DesiredState: cannot read state, err: %v", err)
	}
	for _, s := range state.UsedAccounts {
		if s.ApplicationName == applicationName {
			natsAccount = s.Account
		}
	}

	// Reserve new nats account if possible
	if natsAccount == "" {
		natsAccount, err = c.AllocateNatsAccount(applicationName)
		if err != nil {
			return nil, err
		}
	}

	secrets, err := c.client.CoreV1().Secrets(namespace).List(context.Background(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("natsaccount=%s,natsfunction=users", natsAccount),
	})

	if err != nil {
		return nil, fmt.Errorf("cannot access secrets")
	}

	if len(secrets.Items) == 0 {
		return nil, fmt.Errorf("no secrets found for applicationName %s", applicationName)
	}

	res := map[string]string{}
	for _, user := range usernames {
		res[user] = string(secrets.Items[0].Data[fixedUsername])
	}

	return res, nil
}

// DeleteAccount deletes the account from the credentials.
func (c *CredsSecrets) DeleteAccount(applicationName string) error {
	// check if the account is used.
	// if used delete it
	state, err := ReadState(c.client)
	if err != nil {
		return fmt.Errorf("cannot read state, err: %v", err)
	}
	accountUsed := false
	for i, usedAccount := range state.UsedAccounts {
		if usedAccount.ApplicationName == applicationName {
			state.UsedAccounts = append(state.UsedAccounts[:i], state.UsedAccounts[i+1:]...)
			accountUsed = true
			break
		}
	}
	if !accountUsed {
		return fmt.Errorf("'%s' not used", applicationName)
	}

	err = UpdateState(c.client, state)
	if err != nil {
		return fmt.Errorf("cannot update state, err: %v", err)
	}

	return nil
}
