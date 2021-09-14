package controllers

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/site"
	"github.com/oit-sec-lab/dnt-verify-server/src/interfaces/network"
	usecase "github.com/oit-sec-lab/dnt-verify-server/src/usecase/site"
)

type Controller struct {
	siteInteractor usecase.SiteInteractor
}

func (controller *Controller) VerifyGPC(c Context, s string) {
	var net network.GpcRepository
	sites, err := controller.siteInteractor.FindByURL(s)
	ii := interface{}(sites)
	gg := ii.(string)
	if err != nil {
		sitess, err := controller.siteInteractor.VerifyGPC(s)
		i := interface{}(err)
		g := i.(string)
		if err != nil {
			c.JSON(net.CheckGPC(g))
			return
		}
		u := site.Site{}
		c.Bind(&u)
		err = controller.siteInteractor.Add(u)
		if err != nil {
			a := interface{}(err)
			b := a.(string)
			c.JSON(net.CheckGPC(b))
			return
		}
		d := interface{}(sitess)
		e := d.(string)
		c.JSON(net.CheckGPC(e))
	} else {
		c.JSON(net.CheckGPC(gg))
	}
}
