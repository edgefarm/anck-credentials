/*
Copyright © 2022 Ci4Rail GmbH <engineering@ci4rail.com>
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
	"fmt"

	api "github.com/edgefarm/anck-credentials/pkg/apis/config/v1alpha1"
	common "github.com/edgefarm/anck/pkg/common"
	resources "github.com/edgefarm/anck/pkg/resources"
	"github.com/hsson/once"
)

const (
	// natsServerSecret is the name of the secret that contains the nats server
	natsServerSecretName = "nats-server-infos"
	// natsServerAddressKey is the key of the address of the nats server
	natsServerAddressKey = "NATS_ADDRESS"
	leafServerAddressKey = "LEAF_ADDRESS"
)

var (
	natsAccountInstance *api.ServerInformationResponse
)

// GetNatsServerInfos returns the nats server information
func (c *CredsSecrets) GetNatsServerInfos() (*api.ServerInformationResponse, error) {
	o := once.Error{}
	err := o.Do(func() error {
		sysAccount, err := c.GetSysAccount()
		if err != nil {
			return err
		}
		cont, err := resources.ReadSecret(natsServerSecretName, common.AnckNamespace)
		if err != nil {
			return err
		}
		if _, ok := cont[natsServerAddressKey]; !ok {
			return fmt.Errorf("secret %s does not contain key %s", natsServerSecretName, natsServerAddressKey)
		}
		if _, ok := cont[leafServerAddressKey]; !ok {
			return fmt.Errorf("secret %s does not contain key %s", natsServerSecretName, leafServerAddressKey)
		}

		natsAccountInstance = &api.ServerInformationResponse{
			SysAccount: sysAccount,
			Addresses: &api.Addresses{
				NatsAddress: cont[natsServerAddressKey],
				LeafAddress: cont[leafServerAddressKey],
			},
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	if natsAccountInstance == nil {
		return nil, fmt.Errorf("NatsServer was not initialized properly")
	}

	return natsAccountInstance, nil
}
