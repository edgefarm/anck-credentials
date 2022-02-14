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
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/edgefarm/anck-credentials/pkg/creds"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	api "github.com/edgefarm/anck-credentials/pkg/apis/config/v1alpha1"
)

const (
	namespace      = "anck"
	stateSecret    = "anck-credentials-state"
	fixedUsername  = "user"
	passwortLength = 30
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
func (c *CredsSecrets) DesiredState(accountName string, usernames []string) (*api.DesiredStateResponse, error) {
	var natsAccount = ""

	// Check current state if application name is already used
	state, err := ReadState(c.client)
	if err != nil {
		return nil, fmt.Errorf("DesiredState: cannot read state, err: %v", err)
	}
	for _, s := range state.UsedAccounts {
		if s.ApplicationName == accountName {
			natsAccount = s.Account
		}
	}

	// Reserve new nats account if possible
	if natsAccount == "" {
		natsAccount, err = c.AllocateNatsAccount(accountName)
		if err != nil {
			return nil, err
		}
		c.State, err = ReadState(c.client)
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
		return nil, fmt.Errorf("no secrets found for account %s", accountName)
	}
	accountNameIndex := -1
	configuredUsers := func() []string {
		var users []string
		for i, user := range c.State.UserMappings {
			if user.ApplicationName == accountName {
				accountNameIndex = i
				for _, creds := range user.Credentials {
					users = append(users, creds.Username)
				}
			}
		}
		return users
	}()
	unconfigured, deleted := unconfiguredUsers(configuredUsers, usernames)
	fmt.Println("Unconfigured users: ", unconfigured)
	fmt.Println("Deleted users: ", deleted)
	userCreds := []*api.Credentials{}
	if accountNameIndex != -1 {
		userCreds = c.State.UserMappings[accountNameIndex].Credentials
	}

	for _, user := range unconfigured {
		fmt.Printf("Generating secret for user %s\n", user)
		secret, err := GenerateRandomString(passwortLength)
		if err != nil {
			return nil, fmt.Errorf("cannot generate random string for user %s", user)
		}
		userCreds = append(userCreds, &api.Credentials{
			Username:        user,
			Password:        secret,
			Creds:           string(secrets.Items[0].Data[fixedUsername]),
			UserAccountName: fmt.Sprintf("%s.%s", accountName, user),
		})
	}

	for _, user := range deleted {
		fmt.Printf("Deleting secret for user %s\n", user)
		for i, creds := range userCreds {
			if creds.Username == user {
				userCreds = append(userCreds[:i], userCreds[i+1:]...)
			}
		}
	}

	userMappingIndex := -1
	for i, mapping := range c.State.UserMappings {
		if mapping.ApplicationName == accountName {
			userMappingIndex = i
		}
	}
	if userMappingIndex == -1 {
		c.State.UserMappings = append(c.State.UserMappings, UserMapping{
			ApplicationName: accountName,
			Credentials:     userCreds,
		})
	} else {
		c.State.UserMappings[userMappingIndex].Credentials = userCreds
	}
	err = UpdateState(c.client, c.State)
	if err != nil {
		return nil, fmt.Errorf("cannot update state")
	}

	fmt.Printf("Mapped nats account '%s' to account '%s'\n", natsAccount, accountName)
	res := &api.DesiredStateResponse{
		Creds: userCreds,
		DeletedUserAccountNames: func() []string {
			deletedUserAccountNames := []string{}
			for _, user := range deleted {
				deletedUserAccountNames = append(deletedUserAccountNames, fmt.Sprintf("%s.%s", accountName, user))
			}
			return deletedUserAccountNames
		}(),
	}
	return res, nil
}

// unconfiguredUsers returns two lists:
// 1. a list of users that are not configured yet
// 2. a list of users that are considered as deleted
func unconfiguredUsers(currentlyConfigured []string, userList []string) ([]string, []string) {
	var unconfigured []string
	var deleted []string
	for _, user := range userList {
		if !contains(currentlyConfigured, user) {
			unconfigured = append(unconfigured, user)
		}
	}
	for _, user := range currentlyConfigured {
		if !contains(userList, user) {
			deleted = append(deleted, user)
		}
	}
	return unconfigured, deleted
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// var diff []string
// for i := 0; i < 2; i++ {
// 	for _, s1 := range currentlyConfigured {
// 		found := false
// 		for _, s2 := range users2 {
// 			if s1 == s2 {
// 				found = true
// 				break
// 			}
// 		}
// 		// String not found. We add it to return slice
// 		if !found {
// 			diff = append(diff, s1)
// 		}
// 	}
// 	// Swap the slices, only if it was the first loop
// 	if i == 0 {
// 		currentlyConfigured, users2 = users2, currentlyConfigured
// 	}
// }
// return diff

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

// DeleteAccount deletes the account from the credentials.
func (c *CredsSecrets) DeleteAccount(accountName string) error {
	// check if the account is used.
	// if used delete it
	state, err := ReadState(c.client)
	if err != nil {
		return fmt.Errorf("cannot read state, err: %v", err)
	}
	accountUsed := false
	for i, usedAccount := range state.UsedAccounts {
		if usedAccount.ApplicationName == accountName {
			fmt.Printf("Freeing nats account '%s'\n", usedAccount.Account)
			state.UsedAccounts = append(state.UsedAccounts[:i], state.UsedAccounts[i+1:]...)
			accountUsed = true
			break
		}
	}
	if !accountUsed {
		return fmt.Errorf("'%s' not used", accountName)
	}

	err = UpdateState(c.client, state)
	if err != nil {
		return fmt.Errorf("cannot update state, err: %v", err)
	}

	return nil
}
