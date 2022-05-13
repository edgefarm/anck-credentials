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
	"context"
	"fmt"

	api "github.com/edgefarm/anck-credentials/pkg/apis/config/v1alpha1"
	common "github.com/edgefarm/anck/pkg/common"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *CredsSecrets) GetSysAccount() (*api.SysAccount, error) {
	secrets, err := c.client.CoreV1().Secrets(common.AnckNamespace).List(context.Background(), metav1.ListOptions{
		LabelSelector: "natsfunction=sys-account",
	})
	if err != nil {
		return nil, err
	}

	if len(secrets.Items) == 0 {
		return nil, fmt.Errorf("no sys-account secret found")
	}

	if _, ok := secrets.Items[0].Data["operator-jwt"]; !ok {
		return nil, fmt.Errorf("secret %s does not contain key %s", secrets.Items[0].Name, "operator-jwt")
	}
	if _, ok := secrets.Items[0].Data["sys-public-key"]; !ok {
		return nil, fmt.Errorf("secret %s does not contain key %s", secrets.Items[0].Name, "sys-public-key")
	}
	if _, ok := secrets.Items[0].Data["sys-creds"]; !ok {
		return nil, fmt.Errorf("secret %s does not contain key %s", secrets.Items[0].Name, "sys-creds")
	}
	if _, ok := secrets.Items[0].Data["sys-jwt"]; !ok {
		return nil, fmt.Errorf("secret %s does not contain key %s", secrets.Items[0].Name, "sys-jwt")
	}

	return &api.SysAccount{
		OperatorJWT:  string(secrets.Items[0].Data["operator-jwt"]),
		SysPublicKey: string(secrets.Items[0].Data["sys-public-key"]),
		SysCreds:     string(secrets.Items[0].Data["sys-creds"]),
		SysJWT:       string(secrets.Items[0].Data["sys-jwt"]),
	}, nil
}
