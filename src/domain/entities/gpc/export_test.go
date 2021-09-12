package gpc

import "testing"

func TestGpcGetters(t *testing.T) {
	g := NewGpc(true)
	tests := []struct {
		expected interface{}
		got      interface{}
	}{
		{g.enable, g.Enable()},
	}

	for _, test := range tests {
		if test.expected != test.got {
			t.Fatalf("expected: %v; got: %v\n", test.expected, test.got)
		}
	}
}

