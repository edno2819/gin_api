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

// @BasePath /

// PingExample godoc
// @Summary get users
// @Schemes
// @Description Get all users on db
// @Tags example
// @Produce json
// @Success 200 {"ID":1,"Name":"admin","Password":"123456","UpdatedAt":1686484856995411384,"CreatedAt":1686484856}  AQUI
// @Router /getUsers [get]
func getUsers(ctx *gin.Context) {
	var fist_user models.User
	ctx.MustGet("db").(*gorm.DB).First(&fist_user, "id = ?", 1)
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
