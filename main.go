package main

import (
	"gin_api/app/middlewares"
	"gin_api/app/models"
	"gin_api/app/routers"
	"gin_api/config"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	var conf *config.Config = config.GetConfig()
	config.SetupOutputGin(conf)

	db := models.DatabaseConnection(conf)

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	server.Use(middlewares.DatabaseContext(db))
	server.Use(middlewares.BasicAuth())

	// HTML
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	// Swagger /docs/index.html
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Routers
	routers.BasicRouters(server, "")
	routers.APIRouters(server, "api")

	server.Run(":" + conf.Port)
}
