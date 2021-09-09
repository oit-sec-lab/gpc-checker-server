package repositories

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/site"
)

type ISiteRepository interface {
	Store(site.Site) error
	FindByURL(string) (site.Site, error)
	CheckGPC(string) (bool, error)
	Exists(string) (bool, error)
}
