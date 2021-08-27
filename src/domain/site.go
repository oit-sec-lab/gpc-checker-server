package domain

import "fmt"

const InvalidURL = "invalid empty url string"

type Site struct {
	id  int
	url string
	gpc bool
}

type Sites []Site

func NewSite(i int, u string, g bool) (Site, error) {
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

func (s Site) GPC() bool {
	return s.gpc
}
