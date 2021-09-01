package entities

import (
	"testing"
)

func TestSiteGetters(t *testing.T) {
	s, _ := NewSite(1, "wwww.example.com", true)
	tests := []struct {
		expected interface{}
		got      interface{}
	}{
		{s.id, s.ID()},
		{s.url, s.URL()},
		{s.gpc, s.GPC()},
	}

	for _, test := range tests {
		if test.expected != test.got {
			t.Fatalf("expected: %v; got: %v\n", test.expected, test.got)
		}
	}
}
