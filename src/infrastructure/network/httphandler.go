package network

import (
	interfaceNet "github.com/oit-sec-lab/gpc-checker-server/src/interfaces/network"
	"io/ioutil"
	"net/http"
)

type HttpHandler struct {
	client *http.Client
}

func NewHttpHandler() interfaceNet.HttpHandler {
	httpHandler := new(HttpHandler)
	httpHandler.client = new(http.Client)
	return httpHandler
}


func (handler *HttpHandler) GET(url string) (interfaceNet.Response, error) {
	res := HttpResult{}
	req, e := http.NewRequest(http.MethodGet, url, nil)
	if e != nil {
		return res, e
	}
	r, e := handler.client.Do(req)
	if e != nil {
		return res, e
	}

	a, _ := ioutil.ReadAll(r.Body)
	res.body = string(a)

	defer r.Body.Close()

	res.response = r
	return res, e
}

type HttpResult struct {
	response *http.Response
	body string
}


func (hr HttpResult) Status() string{
	return hr.response.Status
}

func (hr HttpResult) StatusCode() int {
	return hr.response.StatusCode
}

func (hr HttpResult) Protocol() string {
	return hr.response.Proto
}

func (hr HttpResult) Body() string {
	return hr.body
}
