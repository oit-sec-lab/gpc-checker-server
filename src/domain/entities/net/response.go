package net

type Response struct {
	status     string // e.g. "200 OK"
	statusCode int    // e.g. 200
	proto      string // e.g. "HTTP/1.0"
	body string
}

func NewResponse(s string, sc int, p string, body string) Response {
	return Response{s, sc, p, body}
}

func (r Response) Status() string{
	return r.status
}

func (r Response) StatusCode() int {
	return r.statusCode
}

func (r Response) Protocol() string {
	return r.proto
}

func (r Response) Body() string {
	return r.body
}