package infrastructure

import (
    gin "gopkg.in/gin-gonic/gin.v1"
	"interfaces/controllers"
)

var Router *gin.Engine

func init() {
    router := gin.Default()

    siteController := controllers.NewSiteController(NewSqlHandler())

    router.POST("/sites", func(c *gin.Context) {siteController.VerifyGPC(c)})

    Router = router
}
