package controllers

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/interfaces/network"
	siteUsecase "github.com/oit-sec-lab/dnt-verify-server/src/usecase/site"
	"server/interfaces/database"
	gpcUsecase "github.com/oit-sec-lab/dnt-verify-server/src/usecase/site/gpc"
	"fmt"
	"encoding/json"
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

func (controller *Controller) VerifyGPC(c Context)(ret_final string){
	var url_i Entity
	var ret Ret_Struct
	ret_slice := []Ret_Struct{}
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
			ret.Id = url_i.Sites[i].Id
			ret.Url = url_i.Sites[i].Url
			ret.Gpc = sitess.Enable
			ret_slice = append(ret_slice,ret)
			c.JSON(200,ret)
		} else {
			ret.Id = url_i.Sites[i].Id                           
                        ret.Url = sites.URL()                         
                        ret.Gpc = sites.GPC().Enable
			ret_slice = append(ret_slice,ret)
			c.JSON(200,ret)
		}
	}
	ret_final_byte, _ := json.Marshal(ret_slice)
	ret_final = string(ret_final_byte)
	fmt.Println(ret_final)
	return string(ret_final)
}
