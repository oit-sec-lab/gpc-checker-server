package controllers

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/interfaces/network"
	siteUsecase "github.com/oit-sec-lab/dnt-verify-server/src/usecase/site"
	"server/interfaces/database"
	gpcUsecase "github.com/oit-sec-lab/dnt-verify-server/src/usecase/site/gpc"
	"fmt"
)

type Controller struct {
	siteInteractor siteUsecase.SiteInteractor
}

type URLs struct {
	Id int `json:"id"`
	Url string `json:"url"`
}

type Entity struct{
	Sites []*URLs `json:"sites"`
}

func NewSiteController(sqlHandler database.SqlHandler, httpHandler network.HttpHandler) *Controller {
	return &Controller{
		siteInteractor: siteUsecase.NewSiteInteractor(&database.SiteRepository{SqlHandler: sqlHandler},gpcUsecase.NewGpcInteractor(&network.GpcRepository{HttpHandler: httpHandler})),
	}
}

func (controller *Controller) VerifyGPC(c Context) {
	var url_i Entity
	c.Bind(&url_i)
	for i:=0; i<len(url_i.Sites); i++{
		s := url_i.Sites[i].Url
		fmt.Print(s)
		sites, err := controller.siteInteractor.FindByURL(s)
		if err != nil {
			sitess, err := controller.siteInteractor.VerifyGPC(s)
			if err != nil {
				c.JSON(500, err)
				return
			}
			c.JSON(200,sitess)
		} else {
			c.JSON(200,sites.GPC())
		}
	}
}
