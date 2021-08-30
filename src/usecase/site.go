package usecase

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities"
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/repositories"
)

type SiteInteractor struct {
	SiteRepository repositories.ISiteRepository
}

func (interactor *SiteInteractor) Add(s entities.Site) (err error) {
	err = interactor.SiteRepository.Store(s)
	return
}

func (interactor *SiteInteractor) FindByURL(u string) (Site entities.Site, err error) {
	Site, err = interactor.SiteRepository.FindByURL(u)
	return
}

func (interactor *SiteInteractor) VerifyGPC(u string) (gpc bool, err error) {
	Site, err := interactor.SiteRepository.FindByURL(u)
	if err != nil {
		return interactor.SiteRepository.CheckGPC(Site.URL())
	} else {
		Site.GPC()
	}
	return
}
