package network

type HttpHandler interface {
	GET(string) (Response, error)
}

type Response interface {
	Status() string
	StatusCode() int
	Protocol() string
	Body() string
}

