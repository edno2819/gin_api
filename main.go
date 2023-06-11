package main

import (
	"gin_api/config"
	"gin_api/controller"
	"gin_api/middlewares"
	"gin_api/models"
	"gin_api/service"
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

// Organizar as Rotas em um Package
// Criar testes
// Criar CI/CD

func main() {
	var conf *config.Config = config.GetConfig()
	config.SetupOutputGin(conf)

	db := models.DatabaseConnection(conf)

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	server.Use(middlewares.DatabaseContext(db))
	server.Use(middlewares.BasicAuth())

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	// Login Endpoint: Authentication + Token creation
	server.GET("/", func(ctx *gin.Context) {
		var fist_user models.User
		ctx.MustGet("db").(*gorm.DB).First(&fist_user, "id = ?", 1)
		ctx.JSON(http.StatusUnauthorized, fist_user)

	})

	// Login Endpoint: Authentication + Token creation
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

	server.Run(":" + conf.Port)
}
