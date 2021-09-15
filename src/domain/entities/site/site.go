package site

import (
	"fmt"
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/gpc"
)

const InvalidURL = "invalid empty url string"

type Site struct {
	url string
	gpc gpc.Gpc
}

type Sites []Site

func NewSite(u string, g gpc.Gpc) (Site, error) {
	if u == "" {
		return Site{}, fmt.Errorf(InvalidURL)
	}
	return Site{url: u, gpc: g}, nil
}

func (s Site) URL() string {
	return s.url
}

func (s Site) GPC() gpc.Gpc {
	return s.gpc
}
