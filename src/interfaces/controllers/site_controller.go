package controllers

import (
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/site"
	"github.com/oit-sec-lab/dnt-verify-server/src/domain/entities/gpc"
	"github.com/oit-sec-lab/dnt-verify-server/src/interfaces/database"
	"github.com/oit-sec-lab/dnt-verify-server/src/interfaces/net"
	"github.com/oit-sec-lab/dnt-verify-server/src/usecase"
	"strconv"
)

type SiteController struct {
	Interactor usecase.SiteInteractor
}

func init() {

}

func NewSiteController(sqlHandler database.SqlHandler) *SiteController {
	return &SiteController{
		Interactor: usecase.UserInteractor {
			SiteRepository: &database.SiteRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *SiteController) Create(c Context) {
	u := domain.Site{}
	c.Bind(&u)
	site, err := controller.Interactor.Add(u)
	if err != nil {
		return c.JSON(500, NewError(err))
	}
	return c.JSON(201, site)
}

func (controller *SiteController) Index(c Context) {
    sites, err := controller.Interactor.Sites()
    if err != nil {
        return c.JSON(500, NewError(err))
    }
    return c.JSON(200, site)
}

func (controller *SiteController) Show(c Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    site, err := controller.Interactor.SiteById(id)
    if err != nil {
        return c.JSON(500, NewError(err))
    }
    return c.JSON(200, site)
}

func (Interactor *SiteInteractor) VerifyGPC(u string) (gpc bool, err error) {
	return Interactor.JSON(200, VerifyGPC(u))
}

