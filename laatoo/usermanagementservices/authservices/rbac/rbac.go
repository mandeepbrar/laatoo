package ginauth_rbac

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ginauth"
	"github.com/storageutils"
	//"strings"
)

var (
	rb = &rbac{}
)

func init() {
	ginauth.RegisterModule("rbac", rb)
	storageutils.RegisterEventHandler(ginauth.EVENT_JWT_TOKEN_PREPARED, rb.AddToToken)
	storageutils.RegisterEventHandler(ginauth.EVENT_AUTH_COMPLETE, rb.AuthComplete)
}

type rbac struct {
	app        *ginauth.App
	RoleStorer storageutils.Storer
}

func Instance() *rbac {
	return rb
}

func (rb *rbac) Initialize(app *ginauth.App) error {
	app.Logger.Info("rbac: Initializing")
	rb.app = app
	err := rb.configStorer()
	if err != nil {
		return nil
	}
	return nil
}

func (rb *rbac) Serve() error {
	return nil
}

func (rb *rbac) AddToToken(ctx *gin.Context) error {
	userInt, _ := ctx.Get("User")
	user := userInt.(RbacUser)
	tokenInt, _ := ctx.Get("JWT_Token")
	token := tokenInt.(*jwt.Token)
	roles, _ := user.GetRoles()
	token.Claims["Roles"] = roles.ToString()
	user.LoadPermissions(rb.RoleStorer)
	permissions, _ := user.GetPermissions()
	token.Claims["Permissions"] = permissions.ToString()
	return nil
}
func (rb *rbac) AuthComplete(ctx *gin.Context) error {
	userInt, _ := ctx.Get("User")
	user := userInt.(RbacUser)
	tokenInt, _ := ctx.Get("JWT_Token")
	token := tokenInt.(*jwt.Token)
	rolesInt, _ := token.Claims["Roles"]
	user.SetRoles(storageutils.StringToStringSet(rolesInt.(string)))
	permissionsInt, _ := token.Claims["Permissions"]
	user.SetPermissions(storageutils.StringToStringSet(permissionsInt.(string)))
	return nil
}
