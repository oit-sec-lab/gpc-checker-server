package controllers

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/interfaces/network"
	siteUsecase "github.com/oit-sec-lab/dnt-verify-server/src/usecase/site"
	"github.com/oit-sec-lab/dnt-verify-server/src/interfaces/database"
	gpcUsecase "github.com/oit-sec-lab/dnt-verify-server/src/usecase/site/gpc"
	jsonUsecase "github.com/oit-sec-lab/dnt-verify-server/src/usecase/json"
	"net/url"
)

type Controller struct {
	siteInteractor siteUsecase.SiteInteractor
}

func NewSiteController(sqlHandler database.SqlHandler, httpHandler network.HttpHandler) *Controller {
	return &Controller{
		siteInteractor: siteUsecase.NewSiteInteractor(&database.SiteRepository{SqlHandler: sqlHandler},gpcUsecase.NewGpcInteractor(&network.GpcRepository{HttpHandler: httpHandler})),
	}
}

func (controller *Controller) VerifyGPC(c Context)(con Context, err error){
    url_i := jsonUsecase.GenerateURLJsonArray()
    ret_slice := jsonUsecase.GenerateRetJsonArray()
    c.Bind(&url_i)
    for i:=0; i<len(url_i); i++{
        s := jsonUsecase.GetURL(url_i[i])
        u_parse, _ := url.Parse(s)
        s = u_parse.Scheme + "://" + u_parse.Host
        sites, err := controller.siteInteractor.FindByURL(s)
        if err != nil {
            sitess, _ := controller.siteInteractor.VerifyGPC(s)
            if err != nil{
            	return c, err
            }
            ret := jsonUsecase.MakeRetJson(jsonUsecase.GetID(url_i[i]),jsonUsecase.GetURL(url_i[i]),sitess.Enable)
            ret_slice = append(ret_slice,ret)
        } else {
            ret := jsonUsecase.MakeRetJson(jsonUsecase.GetID(url_i[i]),sites.URL(),sites.GPC().Enable)
            ret_slice = append(ret_slice,ret)
        }
    }
    c.JSON(200,ret_slice)
    return
}
