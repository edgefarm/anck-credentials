package config

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	api "github.com/edgefarm/edgefarm.network/pkg/apis/config/v1alpha1"
	"github.com/edgefarm/edgefarm.network/pkg/creds"
)

type Config struct {
	api.UnimplementedConfigServiceServer
	creds.CredsIf
	Port int
	// first string: account
	// second string: username
	// third string: credential for user
	// Credentials map[string]map[string]string
}

func NewConfig(port int, creds creds.CredsIf) *Config {
	return &Config{
		Port:    port,
		CredsIf: creds,
	}
}

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

func (s *Config) DesiredState(ctx context.Context, req *api.DesiredStateRequest) (*api.DesiredStateResponse, error) {
	if req.Account == "" {
		return nil, status.Error(codes.InvalidArgument, "Account cannot be empty")
	}
	for _, username := range req.Username {
		if username == "" {
			return nil, status.Error(codes.InvalidArgument, "Username cannot be empty")
		}
	}

	fmt.Printf("Obtaining secrets for account '%s'\n", req.Account)
	secrets, err := s.CredsIf.DesiredState(req.Account, req.Username)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	res := &api.DesiredStateResponse{
		Credentials: secrets,
	}
	return res, nil
}

func (s *Config) DeleteAccount(context.Context, *api.DeleteAccountRequest) (*api.DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
