package Routes

import (
	"example.com/m/v2/Controllers"
	"example.com/m/v2/Middleware"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	e, err := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	fmt.Println("---------", err)
	v1 := r.Group("api/v1")
	{
		v1.POST("register", Controllers.CreateAUser)
		v1.POST("login", Controllers.Login)
	}
	v1.Use(Middleware.JWTTokenFilter)
	v1.Use(Middleware.NewAuthorizer(e))
	{
		v1.GET("todo", Controllers.GetTodos)
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetATodo)
		v1.PUT("todo/:id", Controllers.UpdateATodo)
		v1.DELETE("todo/:id", Controllers.DeleteATodo)
		v1.GET("users", Controllers.GetAllUser)
		v1.GET("users/:id", Controllers.GetAUser)
	}
	return r
}
