package Middleware

import (
	"example.com/m/v2/Error"
	"example.com/m/v2/Models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func BasicAuth(c *gin.Context) {
	userName, password, hashAuth := c.Request.BasicAuth()

	if hashAuth == true {
		var user Models.User
		err := Models.FindByUserName(&user, userName)
		if err == nil && password == user.PassWord {
			log.WithFields(log.Fields{
				"user": user,
			}).Info("User authenticated")
			return
		}
	}
	message := Error.Message{
		Code:    401,
		Message: "unauthorized",
	}

	c.JSON(http.StatusUnauthorized, message)
	c.AbortWithStatus(http.StatusUnauthorized)
}
