package site

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/gpc"
	"testing"
)

func TestSiteGetters(t *testing.T) {
	s, _ := NewSite("www.example.com", gpc.NewGpc(true))
	tests := []struct {
		expected interface{}
		got      interface{}
	}{
		{s.url, s.URL()},
		{s.gpc, s.GPC()},
	}

	for _, test := range tests {
		if test.expected != test.got {
			t.Fatalf("expected: %v; got: %v\n", test.expected, test.got)
		}
	}
}
