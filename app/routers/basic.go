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

func BasicRouters(server *gin.Engine, endpoint string) {
	server.GET("/", func(ctx *gin.Context) {
		var fist_user models.User
		ctx.MustGet("db").(*gorm.DB).First(&fist_user, "id = ?", 1)
		ctx.JSON(http.StatusUnauthorized, fist_user)

	})

	server.POST("/getToken", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
}
