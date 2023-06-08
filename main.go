package main

import (
	"fmt"
	"gin_api/controller"
	"gin_api/middlewares"
	"gin_api/service"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupOutput() {
	s := fmt.Sprintf("./log/log_%v.log", time.Now().Unix())
	f, _ := os.Create(s)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func main() {

	setupOutput()

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger(), middlewares.BasicAuth())

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll)
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err, video := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusOK, video)
		})
	}

	viewRoutes := server.Group("")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)

	}

	server.Run(":8080")
}
