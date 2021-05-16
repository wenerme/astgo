package agimodels

import (
	"fmt"
	"strconv"
	"strings"
)

type Command interface {
	Command() (string, error)
}
type Response interface {
	Err() error
	Val() (string, error)
}

type Handler interface {
	Command(cmd Command) Response
}
type HandlerFunc func(cmd Command) Response

func (f HandlerFunc) Command(cmd Command) Response {
	return f(cmd)
}

type Client struct {
	Handler
}

func joinCommand(s []interface{}) string {
	sb := strings.Builder{}
	for _, param := range s {
		switch v := param.(type) {
		case string:
			sb.WriteString(v)
		case int:
			sb.WriteString(strconv.Itoa(v))
		case float64:
			sb.WriteString(fmt.Sprint(v))
		case *string:
			if v == nil {
				goto DONE
			}
			sb.WriteString(*v)
		case *int:
			if v == nil {
				goto DONE
			}
			sb.WriteString(strconv.Itoa(*v))
		case *float64:
			if v == nil {
				goto DONE
			}
			sb.WriteString(fmt.Sprint(*v))
		}
		sb.WriteRune(' ')
	}
DONE:
	return strings.TrimSpace(sb.String())
}
