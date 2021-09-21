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

type Ret_Struct struct{
	Id int `json:"id"`
	Url string `json:"url"`
	Gpc bool `json:"gpc"`
}

type Entity struct{
	Sites []*URLs `json:"sites"`
}

func NewSiteController(sqlHandler database.SqlHandler, httpHandler network.HttpHandler) *Controller {
	return &Controller{
		siteInteractor: siteUsecase.NewSiteInteractor(&database.SiteRepository{SqlHandler: sqlHandler},gpcUsecase.NewGpcInteractor(&network.GpcRepository{HttpHandler: httpHandler})),
	}
}

func (controller *Controller) VerifyGPC(c Context){
    var url_i Entity
    var ret Ret_Struct
    ret_slice := []Ret_Struct{}
    c.Bind(&url_i)
    for i:=0; i<len(url_i.Sites); i++{
        s := url_i.Sites[i].Url
        find_sites, err := controller.siteInteractor.FindByURL(s)
        if err != nil {
            verify_sites, _ := controller.siteInteractor.VerifyGPC(s)
	    ret.Id = url_i.Sites[i].Id
            ret.Url = url_i.Sites[i].Url
            ret.Gpc = verify_sites.Enable
            ret_slice = append(ret_slice,ret)
            // c.JSON(200,ret)
        } else {
            ret.Id = url_i.Sites[i].Id
            ret.Url = find_sites.URL()
            ret.Gpc = find_sites.GPC().Enable
            ret_slice = append(ret_slice,ret)
            // c.JSON(200,ret)
        }
    }
    c.JSON(200,ret_slice)
}