package net

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/net"
)

type IWebClient interface {
	GET(string) (net.Response, error)
}
