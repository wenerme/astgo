package ami

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/wenerme/astgo/ami/amimodels"
	"io"
	"sort"
	"strings"
	"testing"
)

func TestMsgIO(t *testing.T) {
	for _, test := range []struct {
		msg *Message
		out string
	}{
		{
			msg: &Message{
				Type: MessageTypeEvent,
				Name: "Name",
				Attributes: map[string]interface{}{
					"Text": "12\r\n34",
					"More": "Yes",
				},
			},
			// more after multi line
			out: "Event: Name\r\nText: 12\r\n34\r\nMore: Yes\r\n\r\n",
		},
		{
			msg: &Message{
				Type: MessageTypeEvent,
				Name: "Name",
				Attributes: map[string]interface{}{
					"Text": "12\r\n34",
				},
			},
			// simple multi line
			out: "Event: Name\r\nText: 12\r\n34\r\n\r\n",
		},

		{
			msg: &Message{
				Type: MessageTypeEvent,
				Name: "FullyBooted",
				Attributes: map[string]interface{}{
					"Uptime": "1234",
				},
			},
			out: "Event: FullyBooted\r\nUptime: 1234\r\n\r\n",
		},
		{msg: MustConvertToMessage(amimodels.FullyBootedEvent{
			Uptime: "1234",
		}),
			out: "Event: FullyBooted\r\nUptime: 1234\r\n\r\n"},
	} {
		msg := test.msg

		// ignore order
		exp := strings.Split(test.out, "\r\n")
		act := strings.Split(msg.Format(), "\r\n")
		sort.Strings(exp)
		sort.Strings(act)
		assert.Equal(t, exp, act)

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
