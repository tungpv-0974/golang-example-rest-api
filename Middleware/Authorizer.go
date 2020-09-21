package Middleware

import (
	"example.com/m/v2/Controllers"
	"example.com/m/v2/Models"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	a:= &BasicAuthorizer{enforcer: e}

	return func(context *gin.Context) {
		if !a.CheckPermission(context.Request) {
			a.RequirePermission(context)
		}
	}
}

func (a *BasicAuthorizer) CheckPermission(r *http.Request) bool {
	userId := Controllers.GetUserIdFromRequest(r)
	var user Models.User
	method := r.Method
	path := r.URL.Path
	err := Models.GetAUser(&user, userId)
	if err != nil {
		return false
	}
	allowed, err := a.enforcer.Enforce(user.Role.Name, path, method)
	if err != nil {
		panic(err)
	}
	return allowed

}

type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(403)
}
