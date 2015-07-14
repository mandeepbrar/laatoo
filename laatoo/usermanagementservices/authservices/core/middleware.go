package ginauth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/storageutils"
)

func (app *App) Authenticate(ctx *gin.Context) {
	headerVal := ctx.Request.Header.Get(AuthToken)
	if headerVal != "" {
		token, err := jwt.Parse(headerVal, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				storageutils.FireEvent(&storageutils.Event{EVENT_AUTH_FAILED, ctx})
				return nil, nil
			}
			return []byte(JWTSecret), nil
		})
		if err == nil && token.Valid {
			userInt := app.UserCreator()
			user := userInt.(storageutils.Storable)
			user.SetId(token.Claims["UserId"].(string))
			ctx.Set("User", userInt)
			ctx.Set("JWT_Token", token)
			storageutils.FireEvent(&storageutils.Event{EVENT_AUTH_COMPLETE, ctx})
			ctx.Next()
			return
		} else {
			ctx.Header(AuthToken, "")
		}
	}
	storageutils.FireEvent(&storageutils.Event{EVENT_AUTH_FAILED, ctx})
}
