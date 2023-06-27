package routers

import (
	"gin_api/app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// postVideos godoc
// @Summary add Video
// @Schemes
// @Description add video in database.
// @Tags example
// @Produce json
// @Success 202 {object} models.Video
// @Router /video [post]
func postVideo(ctx *gin.Context) {
	var video models.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Campos inválidos", "error": err.Error()})
		return
	}
	result := ctx.MustGet("db").(*gorm.DB).Create(&video)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Erro no banco", "error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, video)
}

// getVideos godoc
// @Summary get Videos
// @Schemes
// @Description Get all Videos on db. Authentification Basic is necessary
// @Tags example
// @Produce json
// @Success 200 {array} models.Video
// @Router /video [get]
func getVideo(ctx *gin.Context) {
	var videos []models.Video
	ctx.MustGet("db").(*gorm.DB).Find(&videos)
	ctx.JSON(http.StatusAccepted, videos)
}

// deleteVideoVideos godoc
// @Summary delete Video
// @Schemes
// @Description delete video in database.
// @Tags example
// @Produce json
// @Success 202 "{"Status": "Success"}"
// @Router /video?id=15 [deleteVideo]
func deleteVideo(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "ID inválido"})
		return
	}
	result := ctx.MustGet("db").(*gorm.DB).Delete(&models.Video{}, id)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Fail", "error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"Status": "Success"})
}

func VideoRoute(server *gin.Engine, endpoint string) {
	server.GET(endpoint, getVideo)
	server.POST(endpoint, postVideo)
	server.DELETE(endpoint, deleteVideo)

}
