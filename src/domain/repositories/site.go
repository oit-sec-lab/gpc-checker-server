package repositories

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities"
)

type ISiteRepository interface {
	Store(entities.Site) error
	FindByURL(string) (entities.Site, error)
	CheckGPC(string) (bool, error)
	Exists(string) (bool, error)
}
