package router

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "github.com/oit-sec-lab/gpc-checker-server/src/interfaces/controllers"
    "github.com/oit-sec-lab/gpc-checker-server/src/infrastructure/network"
    "github.com/oit-sec-lab/gpc-checker-server/src/infrastructure/database"
)

var Router *gin.Engine

func init() {
    router := gin.Default()

    router.Use(cors.New(cors.Config{
    	AllowMethods: []string{"*"},
        AllowHeaders: []string{"*"},
        AllowOrigins: []string{"*"},
    }))

    siteController := controllers.NewSiteController(database.NewSqlHandler(),network.NewHttpHandler())

    router.POST("/sites", func(c *gin.Context) {siteController.VerifyGPC(c)})

    Router = router
}
