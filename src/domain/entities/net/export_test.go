package http

import (
	"testing"
)

func TestHttpGetter(t *testing.T) {
	r := NewResponse("200 OK", 200, "HTTP/1.0", "{\n  \"gpc\": true,\n  \"version\": 1\n}")
	tests := []struct {
		expected interface{}
		got      interface{}
	}{
		{r.status, r.Status()},
		{r.statusCode, r.StatusCode()},
		{r.proto, r.Protocol()},
		{r.body, r.Body()},
	}

	for _, test := range tests {
		if test.expected != test.got {
			t.Fatalf("expected: %v; got: %v\n", test.expected, test.got)
		}
	}
}
