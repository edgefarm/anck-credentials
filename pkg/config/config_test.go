package config

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	api "github.com/edgefarm/edgefarm.network/pkg/config/v1alpha1"
	"google.golang.org/grpc"
)

func TestDesiredStateBasic1(t *testing.T) {
	config := NewConfig()
	go config.StartConfigServer()
	cc, err := grpc.Dial(":6000", grpc.WithInsecure())
	assert.Nil(t, err)
	client := api.NewConfigServiceClient(cc)
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	response, err := client.DesiredState(ctx, &api.DesiredStateRequest{
		Account:  "myAccount",
		Username: []string{"user0", "user1"},
	})
	assert.Nil(t, err)
	fmt.Println(response)
	assert.Equal(t, response.Credentials["user0"], "password0")
	assert.Equal(t, response.Credentials["user1"], "password1")
}
