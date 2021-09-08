package usecase

import (
	"fmt"
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/site"
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/repositories"
)

type SiteInteractor struct {
	siteRepository repositories.ISiteRepository
}

const (
	URLNotFound = "url not found"
)

func NewSiteInteractor(r repositories.ISiteRepository) SiteInteractor {
	return SiteInteractor{r}
}
func (interactor *SiteInteractor) Add(s site.Site) (err error) {
	err = interactor.siteRepository.Store(s)
	return
}

func (interactor *SiteInteractor) FindByURL(u string) (Site site.Site, err error) {
	exist, e := interactor.siteRepository.Exists(u)
	if e != nil {
		return site.Site{}, e
	}
	if !exist {
		return site.Site{}, fmt.Errorf(URLNotFound)
	}
	Site, err = interactor.siteRepository.FindByURL(u)
	return
}

func (interactor *SiteInteractor) VerifyGPC(u string) (gpc bool, err error) {
	exist, e := interactor.siteRepository.Exists(u)
	if e != nil {
		return false, e
	}
	if exist {
		s, e := interactor.siteRepository.FindByURL(u)
		if e != nil {
			return false, e
		}
		return s.GPC().Enable(), e
	} else {
		return interactor.siteRepository.CheckGPC(u)
	}
}
