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

	api "github.com/edgefarm/anck-credentials/pkg/apis/config/v1alpha1"
	"github.com/edgefarm/anck-credentials/pkg/creds"
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
	network := req.Network
	// check if account has spaces in it
	if strings.Contains(network, " ") {
		return nil, status.Errorf(codes.InvalidArgument, "Network '%s' contains spaces", network)
	}

	if network == "" {
		return nil, status.Error(codes.InvalidArgument, "Network cannot be empty")
	}

	for _, participant := range req.Participants {
		if participant == "" {
			return nil, status.Error(codes.InvalidArgument, "Participant cannot be empty")
		}
	}

	fmt.Printf("Obtaining secrets for network name '%s'\n", network)
	secrets, err := s.CredsIf.DesiredState(network, req.Participants)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}
	return secrets, nil
}

// DeleteAccount rpc deletes the account from the credentials.
func (s *Config) DeleteNetwork(ctx context.Context, req *api.DeleteNetworkRequest) (*api.DeleteNetworkResponse, error) {
	fmt.Printf("Deleting network '%s'\n", req.Network)
	err := s.CredsIf.DeleteNetwork(req.Network)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot delete network %s", req.Network)
	}
	return &api.DeleteNetworkResponse{}, nil
}

func (s *Config) SysAccount(ctx context.Context, req *api.SysAccountRequest) (*api.SysAccountResponse, error) {
	fmt.Printf("Obtaining sys account request\n")
	res, err := s.CredsIf.SysAccount()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot obtain sys account: %s", err)
	}
	return res, nil

}
