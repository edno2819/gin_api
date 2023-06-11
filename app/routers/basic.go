package routers

import (
	"gin_api/app/controller"
	"gin_api/app/models"
	"gin_api/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	videoService service.VideoService = service.New()
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

// getUsers godoc
// @Summary get users
// @Schemes
// @Description Get all users on db. Authentification Basic is necessary
// @Tags example
// @Produce json
// @Success 200 {array} models.User
// @Router / [get]
func getUsers(ctx *gin.Context) {
	var fist_user []models.User
	ctx.MustGet("db").(*gorm.DB).Find(&fist_user)
	ctx.JSON(http.StatusUnauthorized, fist_user)
}

func getToken(ctx *gin.Context) {
	token := loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, nil)
	}
}

func BasicRouters(server *gin.Engine, endpoint string) {
	server.GET("/", getUsers)
	server.POST("/getToken", getToken)
}
