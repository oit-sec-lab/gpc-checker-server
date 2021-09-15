package router

import (
    "github.com/gin-gonic/gin"
    // "github.com/oit-sec-lab/dnt-verify-server/src/interfaces/controllers"
    "server/interfaces/controllers"
    "server/infrastructure/network"
    "server/infrastructure/database"
)

var Router *gin.Engine

func init() {
    router := gin.Default()

    siteController := controllers.NewSiteController(database.NewSqlHandler(),network.NewHttpHandler())

    router.POST("/sites", func(c *gin.Context) {siteController.VerifyGPC(c)})

    Router = router
}