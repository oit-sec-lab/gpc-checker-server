package site

import (
	"fmt"
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/gpc"
)

const InvalidURL = "invalid empty url string"

type Site struct {
	id  int
	url string
	gpc gpc.Gpc
}

type Sites []Site

func NewSite(i int, u string, g gpc.Gpc) (Site, error) {
	if u == "" {
		return Site{}, fmt.Errorf(InvalidURL)
	}
	return Site{id: i, url: u, gpc: g}, nil
}

func (s Site) ID() int {
	return s.id
}

func (s Site) URL() string {
	return s.url
}

func (s Site) GPC() gpc.Gpc {
	return s.gpc
}
