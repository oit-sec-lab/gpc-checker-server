package controllers

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/interfaces/network"
	siteUsecase "github.com/oit-sec-lab/dnt-verify-server/src/usecase/site"
	"server/interfaces/database"
	gpcUsecase "github.com/oit-sec-lab/dnt-verify-server/src/usecase/site/gpc"
	"fmt"
//	"encoding/json"
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

func NewSiteController(sqlHandler database.SqlHandler, httpHandler network.HttpHandler) *Controller {
	return &Controller{
		siteInteractor: siteUsecase.NewSiteInteractor(&database.SiteRepository{SqlHandler: sqlHandler},gpcUsecase.NewGpcInteractor(&network.GpcRepository{HttpHandler: httpHandler})),
	}
}

func (controller *Controller) VerifyGPC(c Context){
    var url_i []URLs
    var ret Ret_Struct
    ret_slice := []Ret_Struct{}
    c.Bind(&url_i)
    for i:=0; i<len(url_i); i++{
        s := url_i[i].Url
        sites, err := controller.siteInteractor.FindByURL(s)
        if err != nil {
            sitess, _ := controller.siteInteractor.VerifyGPC(s)
            ret.Id = url_i[i].Id
            ret.Url = url_i[i].Url
            ret.Gpc = sitess.Enable
            ret_slice = append(ret_slice,ret)
        } else {
            ret.Id = url_i[i].Id
            ret.Url = sites.URL()
            ret.Gpc = sites.GPC().Enable
            ret_slice = append(ret_slice,ret)
        }
    }
    c.JSON(200,ret_slice)
    return
}