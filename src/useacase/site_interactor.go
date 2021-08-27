package useacase

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain"
)

type SiteInteractor struct {
	SiteRepository SiteRepository
}

func (interactor *SiteInteractor) Add(s domain.Site) (err error) {
	_, err = interactor.SiteRepository.Store(s)
	return
}

func (interactor *SiteInteractor) SiteById(identifier int) (Site domain.Site, err error) {
	Site, err = interactor.SiteRepository.FindByID(identifier)
	return
}