package secrets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnconfiguredChangeNothing(t *testing.T) {
	assert := assert.New(t)
	currentParticipantList := []string{"participanta", "participantb"}
	newParticipantList := []string{"participanta", "participantb"}
	unconfigured, deleted := unconfiguredParticipants(currentParticipantList, newParticipantList)
	assert.Empty(unconfigured)
	assert.Empty(deleted)
}

func TestUnconfiguredParticipantsNewParticipant(t *testing.T) {
	assert := assert.New(t)
	currentParticipantList := []string{"participanta", "participantb"}
	newParticipantList := []string{"participanta", "participantb", "participantc"}
	unconfigured, deleted := unconfiguredParticipants(currentParticipantList, newParticipantList)
	assert.Equal([]string{"participantc"}, unconfigured)
	assert.Empty(deleted)
}

func TestUnconfiguredParticipantsDeleteParticipant(t *testing.T) {
	assert := assert.New(t)
	currentParticipantList := []string{"participanta", "participantb", "participantc"}
	newParticipantList := []string{"participanta", "participantb"}
	unconfigured, deleted := unconfiguredParticipants(currentParticipantList, newParticipantList)
	assert.Equal([]string{"participantc"}, deleted)
	assert.Empty(unconfigured)
}

func TestUnconfiguredParticipantsAddAndDeleteParticipant(t *testing.T) {
	assert := assert.New(t)
	currentParticipantList := []string{"participanta", "participantb", "participantc"}
	newParticipantList := []string{"participanta", "participantb", "participantd"}
	unconfigured, deleted := unconfiguredParticipants(currentParticipantList, newParticipantList)
	assert.Equal([]string{"participantc"}, deleted)
	assert.Equal([]string{"participantd"}, unconfigured)
}
