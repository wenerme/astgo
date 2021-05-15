package agimodels

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuild(t *testing.T) {
	command, err := HangupCommand{ChannelName: "16"}.Command(nil)
	assert.NoError(t, err)
	assert.Equal(t, "HANGUP 16", command)
}
