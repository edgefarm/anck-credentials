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

package config

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	api "github.com/edgefarm/edgefarm.network/pkg/apis/config/v1alpha1"
	"github.com/edgefarm/edgefarm.network/pkg/creds"
)

// Config represents the configuration for the config server
type Config struct {
	api.UnimplementedConfigServiceServer
	creds.CredsIf
	Port int
}

// NewConfig returns a new Config object
func NewConfig(port int, creds creds.CredsIf) *Config {
	return &Config{
		Port:    port,
		CredsIf: creds,
	}
}

// StartConfigServer starts the config grpc server
func (c *Config) StartConfigServer() error {
	lis, err := net.Listen("tcp", "0.0.0.0:"+fmt.Sprintf("%d", c.Port))
	if err != nil {
		log.Fatalf("Error creating listener: %v\n", err)
	}
	s := grpc.NewServer()
	api.RegisterConfigServiceServer(s, c)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error serving grpc server: %v\n:", err)
	}
	return nil
}

// DesiredState rpc constructs the desired state of the credentials
func (s *Config) DesiredState(ctx context.Context, req *api.DesiredStateRequest) (*api.DesiredStateResponse, error) {
	accountName := req.AccountName
	// check if account has spaces in it
	if strings.Contains(accountName, " ") {
		return nil, status.Errorf(codes.InvalidArgument, "AccountName '%s' contains spaces", accountName)
	}

	if accountName == "" {
		return nil, status.Error(codes.InvalidArgument, "AccountName cannot be empty")
	}

	for _, username := range req.Username {
		if username == "" {
			return nil, status.Error(codes.InvalidArgument, "Username cannot be empty")
		}
	}

	fmt.Printf("Obtaining secrets for account name '%s'\n", accountName)
	secrets, err := s.CredsIf.DesiredState(accountName, req.Username)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	return secrets, nil
}

// DeleteAccount rpc deletes the account from the credentials.
func (s *Config) DeleteAccount(ctx context.Context, req *api.DeleteAccountRequest) (*api.DeleteAccountResponse, error) {
	fmt.Printf("Deleting account '%s'\n", req.AccountName)
	err := s.CredsIf.DeleteAccount(req.AccountName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot delete account %s", req.AccountName)
	}
	return &api.DeleteAccountResponse{}, nil
}
