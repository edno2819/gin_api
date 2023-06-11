package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIRouters(server *gin.Engine, endpoint string) {
	apiRoutes := server.Group("/api")

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
