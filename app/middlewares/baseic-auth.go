package middlewares

import (
	"gin_api/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Accounts map[string]string

func BasicAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, password, _ := ctx.Request.BasicAuth()

		var user models.User
		result := ctx.MustGet("db").(*gorm.DB).Where("Name = ? AND Password = ?", username, password).First(&user)

		if result.Error != nil || result.RowsAffected == 0 {
			ctx.Header("WWW-Authenticate", "Basic realm=Authorization Required")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}
