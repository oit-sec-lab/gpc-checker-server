package controllers

import
(
	"domain"
	"interfaces/database"
	"usecase"
	"strconv"
)

type SiteController struct {
	Interactor usecase.SiteInteractor
}

func NewSiteController (sqlHandler database.SqlHandler) *SiteController {
	return &SiteController{
		Interactor: usecase.UserInteractor {
			SiteRepository: &database.SiteRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *SiteController) Create (c Context) {
	u := domain.Site{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201)
}

func (controller *SiteController) Index(c Context) {
    users, err := controller.Interactor.Sites()
    if err != nil {
        c.JSON(500, NewError(err))
        return
    }
    c.JSON(200, sites)
}

func (controller *SiteController) Show(c Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    site, err := controller.Interactor.SiteById(id)
    if err != nil {
        c.JSON(500, NewError(err))
        return
    }
    c.JSON(200, site)
}

