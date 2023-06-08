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
	"github.com/joho/godotenv"
)

var (
	videoService service.VideoService = service.New()
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

var (
	port string
)

func setupOutput() {
	s := fmt.Sprintf("./log/log_%v.log", time.Now().Unix())
	f, _ := os.Create(s)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func init() {
	godotenv.Load(".env")

	if os.Getenv("PORT") == "" {
		port = "8080"
	} else {
		port = os.Getenv("PORT")
	}
}

func main() {

	setupOutput()

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	// Login Endpoint: Authentication + Token creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
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

	server.Run(":" + port)
}
