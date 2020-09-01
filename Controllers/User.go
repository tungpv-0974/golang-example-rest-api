package Controllers

import (
	"example.com/m/v2/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// create a todo
func CreateAUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateAUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//	get todo with id
func GetAUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetAUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
