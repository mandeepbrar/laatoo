package common

//jwt "github.com/dgrijalva/jwt-go"
//"laatoo/sdk/server/auth"
//"laatoo/sdk/server/core"
//"laatoo/sdk/server/errors"
//	"laatoo/sdk/server/log"
//"time"

const (
	CONF_PVTKEYPATH                   = "pvtkey"
	CONF_PUBLICKEYPATH                = "pvtkey"
	CONF_SECURITY_REALM               = "realm"
	CONF_LOGINSERVICE_USERDATASERVICE = "user_data_svc"
)

/*
func completeAuthentication(ctx core.RequestContext, user auth.User, jwtSecret string, authHeader string) (*core.Response, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	rbac, ok := user.(auth.RbacUser)
	if ok {
		roles, _ := rbac.GetRoles()
		permissions, admin := ctx.GetRolePermissions(roles)
		rbac.SetPermissions(permissions)
		token.Claims["Admin"] = admin
	}
	user.SetJWTClaims(token)
	token.Claims["UserId"] = user.GetId()
	//token.Claims["IP"] = ctx.ClientIP()
	token.Claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, errors.RethrowError(ctx, AUTH_ERROR_JWT_CREATION, err)
	}
	return core.NewServiceResponse(core.StatusSuccess, user, map[string]interface{}{authHeader: tokenString}), nil
}
*/
