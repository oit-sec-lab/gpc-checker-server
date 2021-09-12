package net

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/net"
)

type HttpHandler interface {
	GET(string) (net.Response, error)
}
