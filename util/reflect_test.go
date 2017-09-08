package util

import (
	"fmt"
	"github.com/wenerme/astgo/ami"
	"testing"
)

func TestRelName(t *testing.T) {
	s := &ami.OriginateResponseEvent{}
	h := make(map[string]interface{})
	h["Response"] = "Success"
	h["CallerIDNum"] = "123456"
	err := SetStruct(s, h)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", s)
}
