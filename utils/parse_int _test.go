package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInt(t *testing.T) {
	// refs: https://github.com/golang/go/wiki/TableDrivenTests
	var tableTests = []struct {
		name string
		in   string
		out  int
		ok   bool
	}{
		{name: "Should return not ok", in: "", out: 0, ok: false},
		{name: "Should return ok for zero", in: "0", out: 0, ok: true},
		{name: "Should return ok", in: "1", out: 1, ok: true},
	}

	for _, tt := range tableTests {
		tt := tt // see refs about table driven tests
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ok, n := ParseInt(tt.in)
			assert.Equal(t, tt.out, n)
			assert.Equal(t, tt.ok, ok)
		})
	}
}
