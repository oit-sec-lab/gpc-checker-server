package net

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/net"
	"io/ioutil"
	"net/http"
)

type WebClient struct {
	client *http.Client
}

func NewClient() *http.Client {
	return new(http.Client)
}

func (webclient *WebClient) GET(url string) (net.Response, error) {
	req, e := http.NewRequest(http.MethodGet, url, nil)
	if e != nil {
		return net.Response{}, e
	}
	r, e := webclient.client.Do(req)
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return net.Response{}, err
	}
	return net.NewResponse(r.Status, r.StatusCode, r.Proto, string(body)), e
}
