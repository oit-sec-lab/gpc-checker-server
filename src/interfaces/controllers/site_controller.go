package controllers

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain"
	"github.com/oit-sec-lab/dnt-verify-server/src/usecase"
)

type SiteController struct {
	Interactor usecase.SiteInteractor
}

func (controller *SiteController) VerifyGPC(c Context) {
	site, err := controller.Interactor.FindByURL()
	if err != nil {
		site, err = controller.Interactor.VerifyGPC()
		if err != nil {
			c.JSON(200, NewError(err))
			return
		}
		u := domain.Site{}
		c.Bind(&u)
		err := controller.Interactor.Add(u)
		if err != nil {
			c.JSON(500, NewError(err))
			return
		}
		c.JSON(200, site)
	} else {
		c.JSON(200, site)
	}
}
