package agimodels

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuild(t *testing.T) {
	for _, test := range []struct {
		S string
		C Command
	}{
		{S: "HANGUP 16", C: HangupCommand{}.SetChannelName("16")},
		{S: "CONTROL STREAM FILE fn #", C: ControlStreamFileCommand{FileName: "fn", EscapeDigits: "#"}},
		{S: "CONTROL STREAM FILE fn #", C: ControlStreamFileCommand{FileName: "fn", EscapeDigits: "#"}.SetPausechr("Abc")},
	} {
		s, err := test.C.Command()
		assert.NoError(t, err)
		assert.Equal(t, test.S, s)
	}
}
