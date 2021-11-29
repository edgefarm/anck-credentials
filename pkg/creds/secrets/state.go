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
	"encoding/json"
	"fmt"

	api "github.com/edgefarm/edgefarm.network/pkg/apis/config/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// State stores the current state of the credentials (all used accounts)
type NatsAccountMapping struct {
	Account         string `json:"account"`
	ApplicationName string `json:"application"`
}

// State stores information about the current state of the credentials
type UserMapping struct {
	ApplicationName string             `json:"application"`
	Credentials     []*api.Credentials `json:"credentials"`
}

// State stores the current state of the credsmanager
type State struct {
	// UsedAccounts slice of used accounts
	UsedAccounts []NatsAccountMapping `json:"usedaccounts"`
	UserMappings []UserMapping        `json:"usermapping"`
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
