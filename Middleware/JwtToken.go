package Middleware

import (
	"example.com/m/v2/Controllers"
	"example.com/m/v2/Error"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTTokenFilter(c *gin.Context)  {
	err := Controllers.TokenValid(c.Request)
	if err != nil {
		message := Error.Message{
			Code:    401,
			Message: "unauthorized",
		}

		c.JSON(http.StatusUnauthorized, message)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
}
