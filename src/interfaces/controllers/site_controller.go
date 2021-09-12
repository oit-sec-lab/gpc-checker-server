package controllers

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/site"
	"github.com/oit-sec-lab/dnt-verify-server/src/usecase"
)

type SiteController struct {
	Interactor usecase.SiteInteractor
}

func (controller *SiteController) VerifyGPC(c Context, u string) {
	sites, err := controller.Interactor.FindByURL(u)
	if err != nil {
		sites, err = controller.Interactor.VerifyGPC(u)
		if err != nil {
			c.JSON(200, NewError(err))
			return
		}
		u := site.Site{}
		c.Bind(&u)
		err := controller.Interactor.Add(u)
		if err != nil {
			c.JSON(500, NewError(err))
			return
		}
		c.JSON(200, sites)
	} else {
		c.JSON(200, sites)
	}
}
