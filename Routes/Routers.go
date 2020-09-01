package Routes

import (
	"example.com/m/v2/Controllers"
	"example.com/m/v2/Middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("todo", Middleware.BasicAuth, Controllers.GetTodos)
		v1.POST("todo", Middleware.BasicAuth, Controllers.CreateATodo)
		v1.GET("todo/:id", Middleware.BasicAuth, Controllers.GetATodo)
		v1.PUT("todo/:id", Middleware.BasicAuth, Controllers.UpdateATodo)
		v1.DELETE("todo/:id", Middleware.BasicAuth, Controllers.DeleteATodo)

		v1.POST("register", Controllers.CreateAUser)
		v1.GET("users/:id", Middleware.BasicAuth, Controllers.GetAUser)
	}
	return r
}
