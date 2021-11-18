package config

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	api "github.com/edgefarm/edgefarm.network/pkg/config/v1alpha1"
)

type server struct {
	api.UnimplementedConfigServiceServer
}

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) StartConfigServer() error {
	port := "6000"
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("Error creating listener: %v\n", err)
	}
	s := grpc.NewServer()
	api.RegisterConfigServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error serving grpc server: %v\n:", err)
	}
	return nil
}

func (s *server) DesiredState(context.Context, *api.DesiredStateRequest) (*api.DesiredStateResponse, error) {
	res := &api.DesiredStateResponse{
		Credentials: map[string]string{},
	}
	res.Credentials["user0"] = "password0"
	res.Credentials["user1"] = "password1"
	time.Sleep(1 * time.Second)
	return res, nil
}

func (s *server) DeleteAccount(context.Context, *api.DeleteAccountRequest) (*api.DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
