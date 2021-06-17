package ami

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/wenerme/astgo/ami/amimodels"
)

type MessageType string

const (
	MessageTypeAction   MessageType = "Action"
	MessageTypeEvent    MessageType = "Event"
	MessageTypeResponse MessageType = "Response"
)

type Message struct {
	Type       MessageType // type of message
	Name       string      // name of message
	Attributes map[string]interface{}
}

func ReadMessage(m *Message, rd io.Reader) (err error) {
	r := bufio.NewReader(rd)

	m.Attributes = map[string]interface{}{}
	var line string
	line, err = r.ReadString('\n')
	if err != nil {
		return err
	}
RETRY:
	line = strings.TrimSuffix(line, "\r\n")
	if line == "" {
		goto RETRY
	}
	sp := strings.SplitN(line, ":", 2)
	if len(sp) != 2 {
		return errors.Errorf("invalid type line read: %q", line)
	}
	m.Type = MessageType(sp[0])
	m.Name = strings.TrimSpace(sp[1])
	switch m.Type {
	case MessageTypeAction:
	case MessageTypeEvent:
	case MessageTypeResponse:
	default:
		return errors.Errorf("invalid message type: %q", sp[0])
	}

	var stack [][]string
	for {
		// may contain BOM
		line, err = r.ReadString('\n')
		if err != nil {
			return err
		}
		if line == "\r\n" {
			break
		}
		sp = strings.SplitN(line, ":", 2)

		switch {
		case len(sp) == 2 && strings.HasSuffix(line, "\r\n"):
			// valid line
			stack = append(stack, sp)
		case len(stack) == 0 && len(sp) == 2:
			// first line
			stack = append(stack, sp)
		case len(stack) != 0:
			// continue line
			stack = append(stack, []string{"", line})
		}
	}

	var k, v string
	for _, pair := range stack {
		switch {
		case pair[0] != "" && k != "":
			m.Attributes[k] = strings.TrimSuffix(v, "\r\n")
			k = ""
			v = ""
			fallthrough
		case pair[0] != "":
			k = pair[0]
			v = strings.TrimLeft(pair[1], " ")
		case pair[0] == "":
			v += pair[1]
		}
	}
	if k != "" {
		m.Attributes[k] = strings.TrimSuffix(v, "\r\n")
	}
	return
}
func WriteMessage(m *Message, w io.Writer) (err error) {
	wr := bufio.NewWriter(w)
	_, _ = wr.WriteString(fmt.Sprintf("%v: %v\r\n", m.Type, m.Name))
	for k, v := range m.Attributes {
		_, _ = wr.WriteString(fmt.Sprintf("%v: %v\r\n", k, v))
	}
	_, _ = wr.WriteString("\r\n")
	err = wr.Flush()
	return
}

func (m *Message) Read(r *bufio.Reader) (err error) {
	return ReadMessage(m, r)
}
func (m *Message) Write(w io.Writer) (err error) {
	return WriteMessage(m, w)
}

func (m Message) Format() string {
	b := bytes.Buffer{}
	_ = m.Write(&b)
	return b.String()
}

func (m *Message) AttrString(name string) string {
	if m.Attributes == nil {
		return ""
	}
	msg := m.Attributes[name]
	if msg == nil {
		return ""
	}
	return fmt.Sprint(msg)
}
func (m *Message) Message() string {
	return m.AttrString("Message")
}

func (m *Message) Success() bool {
	if m.Type == MessageTypeResponse && m.Name == "Success" {
		return true
	}
	return false
}

func (m *Message) Error() error {
	if m.Type == MessageTypeResponse && m.Name == "Error" {
		msg := m.Message()
		if msg == "" {
			msg = "error response"
		}
		return errors.New(msg)
	}
	return nil
}
func (m *Message) SetAttr(name string, val interface{}) {
	if m.Attributes == nil {
		m.Attributes = make(map[string]interface{})
	}
	m.Attributes[name] = val
}

func MustConvertToMessage(in interface{}) (msg *Message) {
	m, err := ConvertToMessage(in)
	if err != nil {
		panic(err)
	}
	return m
}
func ConvertToMessage(in interface{}) (msg *Message, err error) {
	msg = &Message{}
	switch a := in.(type) {
	case Message:
		return &a, err
	case *Message:
		return a, err
	case amimodels.Action:
		msg.Type = MessageTypeAction
		msg.Name = a.ActionTypeName()
	case amimodels.Event:
		msg.Type = MessageTypeEvent
		msg.Name = a.EventTypeName()
	default:
		return nil, errors.Errorf("invalid type: %T", in)
	}
	m := make(map[string]interface{})
	err = mapstructure.Decode(in, &m)
	// remove zero
	for k, v := range m {
		rm := v == nil

		// NOTE support require tag, prevent remove required empty value ?
		switch v := v.(type) {
		case string:
			rm = v == ""
		case int:
			rm = v == 0
		}
		if rm {
			delete(m, k)
		}
	}
	msg.Attributes = m
	return
}
