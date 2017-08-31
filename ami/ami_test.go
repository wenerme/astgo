package ami

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseHeader(t *testing.T) {
	assert := assert.New(t)

	headers := make(map[string]interface{})
	assert.NoError(parseHeader(`Event: FullyBooted`, headers))
	assert.Equal(headers["Event"], "FullyBooted")
}

func TestBuildHeader(t *testing.T) {
	assert := assert.New(t)

	headers := make(map[string]interface{})
	headers["Event"] = "FullyBooted"
	headers["Uptime"] = 1234
	s := buildHeaders(headers)
	assert.Equal("Event: FullyBooted\r\nUptime: 1234\r\n\r\n", s)
}

func TestBuildAnyCommand(t *testing.T) {
	assert := assert.New(t)

	headers := make(map[string]interface{})
	headers["Event"] = "FullyBooted"
	headers["Uptime"] = 1234
	s := buildHeaders(headers)
	assert.Equal("Event: FullyBooted\r\nUptime: 1234\r\n\r\n", s)

	command, err := buildAnyCommand(LoginAction{
		Username: "abc",
		Secret:   "def",
	})
	assert.NoError(err)
	assert.Equal(command.Headers, map[string]interface{}{
		"Action":   "Login",
		"Username": "abc",
		"Secret":   "def",
	})
	assert.Equal(command.Type, Action)

	assert.Equal("Action: Login\r\nUsername: abc\r\nSecret: def\r\n\r\n",buildHeaders(command.Headers))
}
