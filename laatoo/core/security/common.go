package security

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"laatoo/sdk/auth"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

func hasPermission(ctx core.RequestContext, perm string, rolePermissions map[string]bool) bool {
	if perm == "" {
		return true
	}
	/*bypass, ok := ctx.Get(CONF_SERVICE_AUTHBYPASS)
	if ok && bypass.(bool) {
		log.Logger.Trace(ctx, "Bypassed permission", "perm", perm, "bypass", bypass)
		return true
	}*/
	if ctx.IsAdmin() {
		return true
	}
	usr := ctx.GetUser().(auth.RbacUser)
	roles, _ := usr.GetRoles()
	log.Logger.Trace(ctx, "Checking roles for permission", "perm", perm) //, "bypass", bypass, "roles", roles)
	for _, role := range roles {
		key := fmt.Sprintf("%s#%s", role, perm)
		val, ok := rolePermissions[key]
		if ok {
			return val
		}
	}
	return false
}

func getUserFromToken(ctx core.RequestContext, userCreator core.ObjectCreator, authHeader string, jwtSecret string) (auth.User, bool, error) {
	tokenVal, ok := ctx.GetString(authHeader)
	if ok {
		token, err := jwt.Parse(tokenVal, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_REQUEST)
			}
			return []byte(jwtSecret), nil
		})
		if err == nil && token.Valid {
			userInt, err := userCreator(ctx, nil)
			if err != nil {
				return nil, false, errors.WrapError(ctx, err)
			}
			user, ok := userInt.(auth.RbacUser)
			if !ok {
				return nil, false, errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
			}
			user.LoadJWTClaims(token)
			admin := false
			adminClaim := token.Claims["Admin"]
			if adminClaim != nil {
				adminVal, ok := adminClaim.(bool)
				if ok {
					admin = adminVal
				}
			}
			return user, admin, nil
		}
	}
	return nil, false, nil
}
