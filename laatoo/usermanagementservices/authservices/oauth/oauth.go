package ginauth_oauth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ginauth"
	"github.com/storageutils"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"net/url"
)

func init() {
	oauth := &oAuth{}
	ginauth.RegisterModule("oauth", oauth)
	ginauth.RegisterAuthType("oauth", oauth)
}

type oAuth struct {
	app            *ginauth.App
	authSuccessful gin.HandlerFunc
	authFailure    gin.HandlerFunc
	providers      map[string]OAuthProvider
}

func (oauth *oAuth) Initialize(app *ginauth.App) error {
	oauth.app = app
	oauth.app.Logger.Debug("oAuth: Initializing")
	return nil
}

func (oauth *oAuth) Serve() error {
	return nil
}

func (oauth *oAuth) InitializeType(authStart gin.HandlerFunc, authFailed gin.HandlerFunc, authSuccessful gin.HandlerFunc) error {
	groupRouter := oauth.app.Router.Group("/auth")
	state := storageutils.RandomString(10)
	for name, _ := range oauth.providers {
		groupRouter.GET(fmt.Sprintf("/%s", name), func(ctx *gin.Context) {
			ctx.Set("Provider", name)
			ctx.Set("State", state)
			authStart(ctx)
		})
		groupRouter.GET(fmt.Sprintf("/%scallback", name), func(ctx *gin.Context) {
			ctx.Set("Provider", name)
			ctx.Set("State", state)
			oauth.CompleteAuthentication(ctx)
		})
	}
	oauth.authFailure = authFailed
	oauth.authSuccessful = authSuccessful
	return nil
}

func (oauth *oAuth) ValidateUser(ctx *gin.Context) error {
	oauth.app.Logger.Debug("oAuth: Validating Credentials")
	name, _ := ctx.Get("Provider")
	stateInt, _ := ctx.Get("State")
	provider := oauth.providers[name.(string)]
	config := provider.GetConfig()
	ctx.Redirect(http.StatusTemporaryRedirect, config.AuthCodeURL(stateInt.(string)))
	return nil
}
func (oauth *oAuth) CompleteAuthentication(ctx *gin.Context) error {
	name, _ := ctx.Get("Provider")
	provider := oauth.providers[name.(string)]
	config := provider.GetConfig()
	sentStateInt, _ := ctx.Get("State")
	state := ctx.Param("")
	if state != sentStateInt.(string) {
		oauth.authFailure(ctx)
	}
	code := ctx.Param("code")
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return fmt.Errorf("Could not validate oauth2 code: %v", err)
	}
	endpointProfile := ""
	response, err := http.Get(endpointProfile + "&access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(bits)
	oauth.app.Logger.Info("oAuth: Authentication Successful")
	return nil
}
