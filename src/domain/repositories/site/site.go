package site

import (
	"github.com/oit-sec-lab/gpc-checker-server/src/domain/entities/site"
)

type ISiteRepository interface {
	Store(site.Site) error
	FindByURL(string) (site.Site, error)
	Exists(string) (bool, error)
}
