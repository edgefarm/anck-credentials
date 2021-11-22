package config

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/phayes/freeport"

// 	api "github.com/edgefarm/edgefarm.network/pkg/apis/config/v1alpha1"
// 	"google.golang.org/grpc"
// )

// func TestDesiredStateBasic1(t *testing.T) {
// 	assert := assert.New(t)
// 	port, err := freeport.GetFreePort()
// 	assert.Nil(err)
// 	config := NewConfig(port)
// 	// add some dummy values
// 	config.Credentials["myAccount"] = map[string]string{"user0": "password0", "user1": "password1"}

// 	go func() {
// 		err := config.StartConfigServer()
// 		assert.Nil(err)
// 	}()

// 	cc, err := grpc.Dial(fmt.Sprintf(":%d", config.Port), grpc.WithInsecure())
// 	assert.Nil(err)
// 	client := api.NewConfigServiceClient(cc)
// 	ctx := context.Background()
// 	response, err := client.DesiredState(ctx, &api.DesiredStateRequest{
// 		Account:  "myAccount",
// 		Username: []string{"user0", "user1"},
// 	})
// 	assert.Nil(err)
// 	assert.Equal(response.Credentials["user0"], "password0")
// 	assert.Equal(response.Credentials["user1"], "password1")
// }

// func TestDesiredStateAccountMissingAccount(t *testing.T) {
// 	assert := assert.New(t)
// 	port, err := freeport.GetFreePort()
// 	assert.Nil(err)
// 	config := NewConfig(port)
// 	go func() {
// 		err := config.StartConfigServer()
// 		assert.Nil(err)
// 	}()

// 	cc, err := grpc.Dial(fmt.Sprintf(":%d", config.Port), grpc.WithInsecure())
// 	assert.Nil(err)
// 	client := api.NewConfigServiceClient(cc)
// 	ctx := context.Background()
// 	_, err = client.DesiredState(ctx, &api.DesiredStateRequest{
// 		Account:  "myAccount",
// 		Username: []string{"user0", "user1"},
// 	})
// 	assert.NotNil(err)
// }
