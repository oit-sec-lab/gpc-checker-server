package site_test

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/gpc"
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/site"
	"testing"
)

func TestNewSite(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		gpc      bool
		shouldFail bool
	}{
		{"gpc false", "www.example.com", false, false},
		{"gpc true", "www.example.com", true, false},
		{"gpc false and empty url", "", false, true},
		{"gpc true and empty url", "", true, true},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, e := site.NewSite(tt.url, gpc.NewGpc(tt.gpc))

			if !tt.shouldFail && e != nil {
				t.Fatalf("unexpected error: %v", e)
			} else if tt.shouldFail && e == nil {
				t.Fatalf("expected error, but got nil")
			}
			t.Logf("site: %+v", s)
		})
	}
}

