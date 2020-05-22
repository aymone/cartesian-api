package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	// refs: https://github.com/golang/go/wiki/TableDrivenTests
	var tableTests = []struct {
		name string
		in   int
		out  int
	}{
		{name: "Should not convert", in: 2, out: 2},
		{name: "Should convert", in: -2, out: 2},
		{name: "Should return zero", in: 0, out: 0},
	}

	for _, tt := range tableTests {
		tt := tt // see refs about table driven tests
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.out, Abs(tt.in))
		})
	}
}
