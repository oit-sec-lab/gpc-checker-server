package gpc_test

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/gpc"
	"testing"
)

func TestNewGpc(t *testing.T) {
	tests := []struct {
		expected bool
		gpc      bool
	}{
		{false, false},
		{true, true},
	}

	for _, test := range tests {
		g := gpc.NewGpc(test.gpc)
		if test.expected != g.Enable() {
			t.Fatalf("expected: %v; got: %v\n", test.expected, g.Enable())
		}
	}
}


