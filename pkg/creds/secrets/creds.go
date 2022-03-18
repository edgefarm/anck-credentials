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

	"github.com/edgefarm/anck-credentials/pkg/creds"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	api "github.com/edgefarm/anck-credentials/pkg/apis/config/v1alpha1"
)

const (
	namespace        = "anck"
	stateSecret      = "anck-credentials-state"
	fixedUsername    = "user"
	accountPublicKey = "account-public-key"
	passwortLength   = 30
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
			if s.NatsAccount == secret.Name {
				unusedNatsAccounts, err = removeAccount(unusedNatsAccounts, s.NatsAccount)
				if err != nil {
					return nil, fmt.Errorf("getUnusedNatsAccounts: cannot remove network %s from unusedNatsAccounts, err: %v", s.NatsAccount, err)
				}
			}
		}
	}
	return unusedNatsAccounts, nil
}

// AllocateNatsAccount allocates a new nats account if possible
func (c *CredsSecrets) AllocateNatsAccount(network string) (string, error) {
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
		natsAccount := unusedNatsAccounts[0]
		state.UsedAccounts = append(state.UsedAccounts, NatsAccountMapping{
			Network:     network,
			NatsAccount: natsAccount,
		})
		err := UpdateState(c.client, state)
		if err != nil {
			return "", err
		}
		return natsAccount, nil
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

func hasDuplicates(s []string) bool {
	visited := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		if visited[s[i]] {
			return true
		} else {
			visited[s[i]] = true
		}
	}
	return false
}

// DesiredState constructs the desired state of credentials for a given application name
func (c *CredsSecrets) DesiredState(network string, participants []string) (*api.DesiredStateResponse, error) {
	if hasDuplicates(participants) {
		return nil, fmt.Errorf("DesiredState: participants list cannot contains duplicates")
	}

	networkParticipantsList := []string{}
	for _, p := range participants {
		networkParticipantsList = append(networkParticipantsList, network+"."+p)
	}

	var natsAccount = ""

	// Check current state if application name is already used
	state, err := ReadState(c.client)
	if err != nil {
		return nil, fmt.Errorf("DesiredState: cannot read state, err: %v", err)
	}
	for _, s := range state.UsedAccounts {
		if s.Network == network {
			natsAccount = s.NatsAccount
		}
	}

	// Reserve new nats account if possible
	if natsAccount == "" {
		natsAccount, err = c.AllocateNatsAccount(network)
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
		return nil, fmt.Errorf("no secrets found for network %s", network)
	}
	networkParticipantIndex := -1
	configuredParticipants := func() []string {
		var participants []string
		for i, participant := range c.State.ParticipantMappings {
			if participant.Network == network {
				networkParticipantIndex = i
				for _, creds := range participant.Credentials {
					participants = append(participants, creds.NetworkParticipant)
				}
			}
		}
		return participants
	}()
	unconfigured, deleted := unconfiguredParticipants(configuredParticipants, networkParticipantsList)
	fmt.Println("Unconfigured participants: ", unconfigured)
	fmt.Println("Deleted participants: ", deleted)
	participantCreds := []*api.Credentials{}
	if networkParticipantIndex != -1 {
		participantCreds = c.State.ParticipantMappings[networkParticipantIndex].Credentials
	}

	for _, participant := range unconfigured {
		fmt.Printf("Generating secret for %s\n", participant)
		participantCreds = append(participantCreds, &api.Credentials{
			Creds:              string(secrets.Items[0].Data[fixedUsername]),
			AccoutPublicKey:    string(secrets.Items[0].Data[accountPublicKey]),
			NetworkParticipant: participant,
		})
	}

	for _, participant := range deleted {
		fmt.Printf("Deleting secret for %s\n", participant)
		for i, creds := range participantCreds {
			if creds.NetworkParticipant == participant {
				participantCreds = append(participantCreds[:i], participantCreds[i+1:]...)
			}
		}
	}

	userMappingIndex := -1
	for i, mapping := range c.State.ParticipantMappings {
		if mapping.Network == network {
			userMappingIndex = i
		}
	}
	if userMappingIndex == -1 {
		c.State.ParticipantMappings = append(c.State.ParticipantMappings, ParticipantMapping{
			Network:     network,
			Credentials: participantCreds,
		})
	} else {
		c.State.ParticipantMappings[userMappingIndex].Credentials = participantCreds
	}
	err = UpdateState(c.client, c.State)
	if err != nil {
		return nil, fmt.Errorf("cannot update state")
	}

	fmt.Printf("Mapped nats account '%s' to network '%s'\n", natsAccount, network)
	res := &api.DesiredStateResponse{
		Creds: participantCreds,
		DeletedParticipants: func() []string {
			deletedNetworkParticipants := []string{}
			deletedNetworkParticipants = append(deletedNetworkParticipants, deleted...)
			return deletedNetworkParticipants
		}(),
	}
	return res, nil
}

// unconfiguredParticipants returns two lists:
// 1. a list of participants that are not configured yet
// 2. a list of participants that are considered as deleted
func unconfiguredParticipants(currentlyConfigured []string, userList []string) ([]string, []string) {
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

// DeleteNetwork deletes the network from the credentials.
func (c *CredsSecrets) DeleteNetwork(network string) error {
	// check if the account is used.
	// if used delete it
	state, err := ReadState(c.client)
	if err != nil {
		return fmt.Errorf("cannot read state, err: %v", err)
	}
	networkUsed := false
	for i, usedAccount := range state.UsedAccounts {
		if usedAccount.Network == network {
			fmt.Printf("Freeing network '%s'\n", usedAccount.Network)
			state.UsedAccounts = append(state.UsedAccounts[:i], state.UsedAccounts[i+1:]...)
			networkUsed = true
			break
		}
	}

	for i, mapping := range state.ParticipantMappings {
		if mapping.Network == network {
			state.ParticipantMappings = append(state.ParticipantMappings[:i], state.ParticipantMappings[i+1:]...)
			break
		}
	}

	if !networkUsed {
		return fmt.Errorf("'%s' not used", network)
	}

	err = UpdateState(c.client, state)
	if err != nil {
		return fmt.Errorf("cannot update state, err: %v", err)
	}

	return nil
}
