package ami

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	amimodels "github.com/wenerme/astgo/pkg/ami/models"
	"io"
	"testing"
)

func TestMsgIO(t *testing.T) {
	for _, test := range []struct {
		msg *Message
		out string
	}{
		{msg: &Message{
			Type: MessageTypeEvent,
			Name: "FullyBooted",
			Attributes: map[string]interface{}{
				"Uptime": "1234",
			},
		},
			out: "Event: FullyBooted\r\nUptime: 1234\r\n\r\n"},
		{msg: MustConvertToMessage(amimodels.FullyBootedEvent{
			Uptime: "1234",
		}),
			out: "Event: FullyBooted\r\nUptime: 1234\r\n\r\n"},
	} {
		msg := test.msg
		assert.Equal(t, test.out, msg.Format())
		r := bufio.NewReader(bytes.NewReader([]byte(test.out)))
		rm := &Message{}
		assert.NoError(t, rm.Read(r))
		assert.Equal(t, msg.Type, rm.Type)
		assert.Equal(t, msg.Name, rm.Name)
		assert.Equal(t, msg.Attributes, rm.Attributes)
		_, err := r.ReadByte()
		assert.Equal(t, io.EOF, err)
	}
}
