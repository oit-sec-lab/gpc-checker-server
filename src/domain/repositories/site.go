package repositories

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities"
)

type SiteRepository interface {
	Store(entities.Site) (int, error)
	FindByID(int) (entities.Site, error)
}
