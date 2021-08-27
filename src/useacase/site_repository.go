package useacase

import "github.com/oit-sec-lab/dnt-verify-server/src/domain"

type SiteRepository interface {
	Store(domain.Site) (int, error)
	FindByID(int) (domain.Site, error)
}
