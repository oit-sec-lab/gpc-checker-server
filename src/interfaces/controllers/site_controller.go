package controllers

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/site"
	"github.com/oit-sec-lab/dnt-verify-server/src/interfaces/network"
	siteUsecase "github.com/oit-sec-lab/dnt-verify-server/src/usecase/site"
	"server/interfaces/database"
	// repo "github.com/oit-sec-lab/dnt-verify-server/src/domain/repositories/site"
	gpcUsecase "github.com/oit-sec-lab/dnt-verify-server/src/usecase/site/gpc"
	// "server/interfaces/network"
)

type Controller struct {
	siteInteractor siteUsecase.SiteInteractor
}

func NewSiteController(sqlHandler database.SqlHandler, httpHandler network.HttpHandler) *Controller {
	return &Controller{
		siteInteractor: siteUsecase.NewSiteInteractor(&database.SiteRepository{SqlHandler: sqlHandler},gpcUsecase.NewGpcInteractor(&network.GpcRepository{HttpHandler: httpHandler})),
	}
}

func (controller *Controller) VerifyGPC(c Context) {
	// s := c.Params("url")
	s := "https://example.com/"
	sites, err := controller.siteInteractor.FindByURL(s)
	if err != nil {
		sitess, err := controller.siteInteractor.VerifyGPC(s)
		if err != nil {
			c.JSON(500,err)
			return
		}
		u := site.Site{}
		c.Bind(&u)
		err = controller.siteInteractor.Add(u)
		if err != nil {
			c.JSON(500,err)
			return
		}
		c.JSON(200,sitess)
	} else {
		c.JSON(200,sites)
	}
}
