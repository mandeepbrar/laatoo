package ginauth_rbac

import (
	"github.com/gin-gonic/gin"
	//"strings"
)

func (rb *rbac) AllowPermissions(permissions []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userInt, _ := ctx.Get("User")
		user := userInt.(RbacUser)
		userPermissions, _ := user.GetPermissions()
		rb.app.Logger.Info(userPermissions.ToString())
		ctx.Next()
	}
}

func (rb *rbac) AllowRoles(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userInt, _ := ctx.Get("User")
		user := userInt.(RbacUser)
		userRoles, _ := user.GetRoles()
		rb.app.Logger.Info(userRoles.ToString())
		ctx.Next()
	}
}
