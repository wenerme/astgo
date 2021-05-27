package agi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseResponse(t *testing.T) {
	for _, test := range []struct {
		L string
		R Response
	}{
		{L: "200 result=1", R: Response{Status: 200, Result: 1, ResultString: "1"}},
		{L: "200 result=0 endpos=8640", R: Response{Status: 200, Result: 0, ResultString: "0", Value: "endpos=8640"}},
	} {
		assert.Equal(t, &test.R, ParseResponse(test.L))
	}
}
