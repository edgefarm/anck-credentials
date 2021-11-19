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
	// first string: account
	// second string: username
	// third string: credential for user
	// Credentials map[string]map[string]string
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
	account := req.Account
	// check if account has spaces in it
	if strings.Contains(account, " ") {
		return nil, status.Errorf(codes.InvalidArgument, "Account '%s' contains spaces", account)
	}

	if account == "" {
		return nil, status.Error(codes.InvalidArgument, "Account cannot be empty")
	}

	for _, username := range req.Username {
		if username == "" {
			return nil, status.Error(codes.InvalidArgument, "Username cannot be empty")
		}
	}

	fmt.Printf("Obtaining secrets for account '%s'\n", account)
	secrets, err := s.CredsIf.DesiredState(account, req.Username)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	res := &api.DesiredStateResponse{
		Credentials: secrets,
	}
	return res, nil
}

// DeleteAccount rpc deletes the account from the credentials.
func (s *Config) DeleteAccount(ctx context.Context, req *api.DeleteAccountRequest) (*api.DeleteAccountResponse, error) {
	err := s.CredsIf.DeleteAccount(req.Account)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot delete account %s", req.Account)
	}
	return &api.DeleteAccountResponse{}, nil
}
