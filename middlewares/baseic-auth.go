package middlewares

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth((gin.Accounts{
		"admin": "123456",
		"pedro": "123456",
	}))
}
