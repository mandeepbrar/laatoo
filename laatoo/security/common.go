package security

import (
	jwt "github.com/dgrijalva/jwt-go"
	"laatoo/sdk/auth"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	//	"laatoo/sdk/log"
	"time"
)

func completeAuthentication(ctx core.RequestContext, user auth.User) (*core.ServiceResponse, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	rbac, ok := user.(auth.RbacUser)
	if ok {
		roles, _ := rbac.GetRoles()
		permissions, admin := ctx.GetRolePermissions(roles)
		rbac.SetPermissions(permissions)
		token.Claims["Admin"] = admin
		ctx.SetAdmin(true)
	}
	user.SetJWTClaims(token)
	token.Claims["UserId"] = user.GetId()
	//token.Claims["IP"] = ctx.ClientIP()
	token.Claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
	ctx.Set("JWT_Token", token)
	jwtSecret := ctx.GetServerVariable(core.JWTSECRETKEY)
	tokenString, err := token.SignedString([]byte(jwtSecret.(string)))
	if err != nil {
		return nil, errors.RethrowError(ctx, AUTH_ERROR_JWT_CREATION, err)
	}
	authHeader := ctx.GetServerVariable(core.AUTHHEADER)
	return core.NewServiceResponse(core.StatusSuccess, user, map[string]interface{}{authHeader.(string): tokenString}), nil
}
