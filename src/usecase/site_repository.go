package usecase

import "../domain"

type SiteRepository interface {
	Store(domain.Site) (int, error)
	FindByID(int) (domain.Site, error)
}
