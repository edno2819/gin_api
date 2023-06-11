package controller

import (
	"gin_api/app/models"
	"gin_api/app/service"
	"gin_api/app/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func init() {
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	validate.RegisterValidation("is-long", validators.ValidateLong)
	validate.RegisterValidation("is-legal", validators.ValidateLegalTitle)
}

type VideoController interface {
	FindAll() []models.Video
	Save(ctx *gin.Context) (error, models.Video)
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func (c *controller) Save(ctx *gin.Context) (error, models.Video) {
	var video models.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err, video
	}
	err = validate.Struct(video)
	if err != nil {
		return err, video
	}
	c.service.Save(video)
	return nil, video
}

func (c *controller) FindAll() []models.Video {
	return c.service.FindAll()
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func New(service service.VideoService) *controller {
	return &controller{
		service: service,
	}
}
