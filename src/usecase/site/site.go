package site

import (
	"fmt"
	entitiesGpc "github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/gpc"
	entitiesSite "github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/site"
	siteRepository "github.com/oit-sec-lab/dnt-verify-server/src/domain/repositories/site"
	gpcInteractor "github.com/oit-sec-lab/dnt-verify-server/src/usecase/site/gpc"
)

type SiteInteractor struct {
	siteRepository siteRepository.ISiteRepository
	gpcInteractor  gpcInteractor.GpcInteractor
}

const (
	URLNotFound = "url not found"
)

func NewSiteInteractor(sr siteRepository.ISiteRepository, gi gpcInteractor.GpcInteractor) SiteInteractor {
	return SiteInteractor{sr, gi}
}
func (interactor *SiteInteractor) Add(s entitiesSite.Site) (err error) {
	err = interactor.siteRepository.Store(s)
	return
}

func (interactor *SiteInteractor) FindByURL(u string) (Site entitiesSite.Site, err error) {
	exist, e := interactor.siteRepository.Exists(u)
	if e != nil {
		return entitiesSite.Site{}, e
	}
	if !exist {
		return entitiesSite.Site{}, fmt.Errorf(URLNotFound)
	}
	Site, err = interactor.siteRepository.FindByURL(u)
	return
}

func (interactor *SiteInteractor) VerifyGPC(u string) (gpc entitiesGpc.Gpc, err error) {
	exist, e := interactor.siteRepository.Exists(u)
	if e != nil {
		return entitiesGpc.Gpc{}, e
	}
	if exist {
		s, e := interactor.siteRepository.FindByURL(u)
		if e != nil {
			return entitiesGpc.Gpc{}, e
		}
		return s.GPC(), e
	} else {
		g, err := interactor.gpcInteractor.CheckGPC(u)
		s, e := entitiesSite.NewSite(u, g)
		if e != nil {
			return entitiesGpc.Gpc{}, e
		}
		e = interactor.Add(s)
		if e != nil {
			return entitiesGpc.Gpc{}, e
		}
		return g, err
	}
}
