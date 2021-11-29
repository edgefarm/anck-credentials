package secrets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnconfiguredChangeNothing(t *testing.T) {
	assert := assert.New(t)
	currentUserList := []string{"usera", "userb"}
	newUserList := []string{"usera", "userb"}
	unconfigured, deleted := unconfiguredUsers(currentUserList, newUserList)
	assert.Empty(unconfigured)
	assert.Empty(deleted)
}

func TestUnconfiguredUsersNewUser(t *testing.T) {
	assert := assert.New(t)
	currentUserList := []string{"usera", "userb"}
	newUserList := []string{"usera", "userb", "userc"}
	unconfigured, deleted := unconfiguredUsers(currentUserList, newUserList)
	assert.Equal([]string{"userc"}, unconfigured)
	assert.Empty(deleted)
}

func TestUnconfiguredUsersDeleteUser(t *testing.T) {
	assert := assert.New(t)
	currentUserList := []string{"usera", "userb", "userc"}
	newUserList := []string{"usera", "userb"}
	unconfigured, deleted := unconfiguredUsers(currentUserList, newUserList)
	assert.Equal([]string{"userc"}, deleted)
	assert.Empty(unconfigured)
}

func TestUnconfiguredUsersAddAndDeleteUser(t *testing.T) {
	assert := assert.New(t)
	currentUserList := []string{"usera", "userb", "userc"}
	newUserList := []string{"usera", "userb", "userd"}
	unconfigured, deleted := unconfiguredUsers(currentUserList, newUserList)
	assert.Equal([]string{"userc"}, deleted)
	assert.Equal([]string{"userd"}, unconfigured)
}
