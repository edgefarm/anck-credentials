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

package creds

import (
	api "github.com/edgefarm/anck-credentials/pkg/apis/config/v1alpha1"
)

// If is an interface to handle the different sources of credentials.
type CredsIf interface {
	DesiredState(network string, participants []string) (*api.DesiredStateResponse, error)
	DeleteNetwork(network string) error
	GetNatsServerInfos() (*api.ServerInformationResponse, error)
}

// Creds contains everything a Credentials uses
type Creds struct {
	Credentials map[string]map[string]string
}
