package security

//jwt "github.com/dgrijalva/jwt-go"
//"laatoo/sdk/auth"
//"laatoo/sdk/core"
//"laatoo/sdk/errors"
//	"laatoo/sdk/log"
//"time"

const (
	CONF_PVTKEYPATH     = "pvtkey"
	CONF_PUBLICKEYPATH  = "pvtkey"
	CONF_SECURITY_REALM = "realm"
)

/*
func completeAuthentication(ctx core.RequestContext, user auth.User, jwtSecret string, authHeader string) (*core.ServiceResponse, error) {
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
