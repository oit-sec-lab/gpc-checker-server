package network

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHttpHandler(t *testing.T) {
	client := NewHttpHandler()
	if client == nil {
		t.Fatalf("unexpected error. client is empty")
	}
}

type Response struct {
	path, query, contenttype, body string
}

func TestHttpHandler_GET(t *testing.T) {
	response := &Response{
		path:        "/.well-known/gpc.json",
		contenttype: "application/json",
		body: `{"gpc": true, "version": 1}`,
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if g, w := r.URL.Path, response.path; g != w {
			t.Errorf("request got path %s, want %s", g, w)
		}
		w.Header().Set("Content-Type", response.contenttype)
		io.WriteString(w, response.body)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewHttpHandler()
	res, e := client.GET(server.URL + "/.well-known/gpc.json")
	if e != nil {
		t.Fatal(e)
	}

	if res.StatusCode() != 200 {
		t.Errorf("expected status code %v, but got status code %v", 200, res.StatusCode())
	}

}
