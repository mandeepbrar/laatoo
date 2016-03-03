package laatooauthentication

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"laatoocore"
	"laatoosdk/auth"
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/utils"
	"net/http"
)

const (
	CONF_AUTHSERVICE_OAUTHPATH_SITES = "oauthsites"
	//login path to be used for local and oauth authentication
	CONF_AUTHSERVICE_OAUTH_AUTHCALLBACK   = "callbackurl"
	CONF_AUTHSERVICE_OAUTH_AUTHURL        = "authurl"
	CONF_AUTHSERVICE_OAUTH_LOGINURL       = "oauthlogin"
	CONF_AUTHSERVICE_OAUTH_LOGININTERCEPT = "intercept"
	CONF_AUTHSERVICE_OAUTH_PROFILEURL     = "profileurl"
	CONF_AUTHSERVICE_OAUTH_TYPE           = "sitetype"
	CONF_AUTHSERVICE_OAUTH_CLIENTID       = "clientid"
	CONF_AUTHSERVICE_OAUTH_CLIENTSECRET   = "clientsecret"
)

type OAuthSite struct {
	sitetype      string
	systemAuthURL string
	callbackURL   string
	profileURL    string
	interceptor   bool
	config        *oauth2.Config
}

type OAuthType struct {
	sites []*OAuthSite
	//method called in case of callback
	authCallback core.HandlerFunc
	//reference to the main auth service
	securityService *SecurityService
}

//method called for creating new auth type
func NewOAuth(ctx core.Context, conf map[string]interface{}, svc *SecurityService) (*OAuthType, error) {
	//create the new auth type
	oauth := &OAuthType{}
	//store the reference to the parent
	oauth.securityService = svc
	log.Logger.Debug(nil, LOGGING_CONTEXT, "OAuthType: Initializing")
	sitesInt, ok := conf[CONF_AUTHSERVICE_OAUTHPATH_SITES]
	if ok {
		sitesMap, ok := sitesInt.(map[string]interface{})
		if ok {
			oauth.sites = make([]*OAuthSite, len(sitesMap))
			i := 0
			for _, v := range sitesMap {
				siteConf := v.(map[string]interface{})
				siteTypeInt, ok := siteConf[CONF_AUTHSERVICE_OAUTH_TYPE]
				if !ok {
					return nil, errors.ThrowError(nil, AUTH_ERROR_OAUTH_MISSING_TYPE)
				}
				siteType := siteTypeInt.(string)
				var endpoint oauth2.Endpoint
				switch siteType {
				case "google":
					endpoint = google.Endpoint
				case "facebook":
					endpoint = facebook.Endpoint
				default:
					return nil, errors.ThrowError(nil, AUTH_ERROR_OAUTH_MISSING_TYPE)
				}
				systemAuthUrlInt, ok := siteConf[CONF_AUTHSERVICE_OAUTH_AUTHURL]
				if !ok {
					return nil, errors.ThrowError(nil, AUTH_ERROR_OAUTH_MISSING_AUTHURL)
				}
				clientIdInt, ok := siteConf[CONF_AUTHSERVICE_OAUTH_CLIENTID]
				if !ok {
					return nil, errors.ThrowError(nil, AUTH_ERROR_OAUTH_MISSING_CLIENTID)
				}
				clientSecretInt, ok := siteConf[CONF_AUTHSERVICE_OAUTH_CLIENTSECRET]
				if !ok {
					return nil, errors.ThrowError(nil, AUTH_ERROR_OAUTH_MISSING_CLIENTSECRET)
				}
				profileInt, ok := siteConf[CONF_AUTHSERVICE_OAUTH_PROFILEURL]
				if !ok {
					return nil, errors.ThrowError(nil, AUTH_ERROR_OAUTH_MISSING_PROFILEURL)
				}
				callbackURLInt, ok := siteConf[CONF_AUTHSERVICE_OAUTH_AUTHCALLBACK]
				if !ok {
					return nil, errors.ThrowError(nil, AUTH_ERROR_OAUTH_MISSING_CALLBACKURL)
				}

				conf := &oauth2.Config{
					ClientID:     clientIdInt.(string),
					ClientSecret: clientSecretInt.(string),
					RedirectURL:  callbackURLInt.(string),
					Scopes:       []string{"openid", "profile", "email"},
					Endpoint:     endpoint,
				}
				interceptor := true
				interceptInt, ok := siteConf[CONF_AUTHSERVICE_OAUTH_LOGININTERCEPT]
				if ok {
					interceptor = (interceptInt.(string) != "false")
				}
				oauth.sites[i] = &OAuthSite{sitetype: siteType, interceptor: interceptor, profileURL: profileInt.(string), systemAuthURL: systemAuthUrlInt.(string), callbackURL: callbackURLInt.(string), config: conf}
				i++
			}
		}
	}

	return oauth, nil
}

//initialize auth type called by base auth for initializing
func (oauth *OAuthType) InitializeType(ctx core.Context, authStart core.HandlerFunc, authCallback core.HandlerFunc) error {
	oauth.authCallback = authCallback
	state := utils.RandomString(10)
	for _, site := range oauth.sites {
		log.Logger.Debug(ctx, LOGGING_CONTEXT, "OAuthType: Setting up site", "site", site)
		oauth.securityService.Router.Get(ctx, site.systemAuthURL, nil, func(ctx core.Context) error {
			ctx.Set("Site", site)
			ctx.Set("State", state)
			return authStart(ctx)
		})
		oauth.securityService.Router.Get(ctx, site.systemAuthURL+"/callback", nil, func(ctx core.Context) error {
			ctx.Set("Site", site)
			ctx.Set("State", state)
			if site.interceptor {
				return oauth.InterceptorPage(ctx)
			} else {
				return authCallback(ctx)
			}
		})
		if site.interceptor {
			oauth.securityService.Router.Post(ctx, site.systemAuthURL, nil, func(ctx core.Context) error {
				ctx.Set("Site", site)
				ctx.Set("State", state)
				return authCallback(ctx)
			})
		}
	}

	return nil
}

//validate the local user
//derive the data from context object
func (oauth *OAuthType) ValidateUser(ctx core.Context) error {
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "OAuthProvider: Validating Credentials")

	siteInt := ctx.Get("Site")
	site, _ := siteInt.(*OAuthSite)
	stateInt := ctx.Get("State")
	url := site.config.AuthCodeURL(stateInt.(string))
	ctx.Redirect(http.StatusTemporaryRedirect, url)

	/*//create the user
	usrInt, err := localauth.securityService.CreateUser()
	if err != nil {
		return errors.RethrowHttpError(laatoocore.AUTH_ERROR_USEROBJECT_NOT_CREATED, ctx, err)
	}

	//ctx.Request().Body
	err = ctx.Bind(usrInt)
	if err != nil {
		return errors.RethrowHttpError(AUTH_ERROR_INCORRECT_REQ_FORMAT, ctx, err)
	}

	//get the ide of the user to be tested
	usr := usrInt.(auth.LocalAuthUser)
	id := usr.GetId()

	//get the tested user from database
	testedUser, err := localauth.securityService.GetUserById(id)
	if err != nil {
		return errors.RethrowHttpError(AUTH_ERROR_USER_NOT_FOUND, ctx, err)
	}
	if testedUser == nil {
		return errors.ThrowHttpError(AUTH_ERROR_USER_NOT_FOUND, ctx)
	}

	//compare the user requested with the user from database
	existingUser := testedUser.(auth.LocalAuthUser)
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.GetPassword()), []byte(usr.GetPassword()))
	existingUser.SetPassword("")
	if err != nil {
		return errors.RethrowHttpError(AUTH_ERROR_WRONG_PASSWORD, ctx, err)
	} else {
		existingUser.SetPassword("")
		ctx.Set("User", testedUser)
		return localauth.authCallback(ctx)
	}*/
	return nil
}

func (oauth *OAuthType) GetName() string {
	return "oauth"
}

//complete authentication
func (oauth *OAuthType) InterceptorPage(ctx core.Context) error {
	siteInt := ctx.Get("Site")
	site, _ := siteInt.(*OAuthSite)
	sentStateInt := ctx.Get("State")
	state := ctx.Query("state")
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "OAuthType: Received code", "state", state)
	if state != sentStateInt.(string) {
		log.Logger.Debug(ctx, LOGGING_CONTEXT, "OAuthType: State mismatch", "state", state, "sentStateInt", sentStateInt)
		return errors.ThrowError(ctx, AUTH_ERROR_USER_NOT_FOUND)
	}
	code := ctx.Query("code")
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "OAuthType: Received code", "code", code)
	return ctx.HTML(http.StatusOK, "<html><body onload='var data = {type:\"%s\",state:\"%s\", code:\"%s\"}; window.opener.postMessage(data, \"*\"); window.close();'></body></html", site.sitetype, state, code)
}

//complete authentication
func (oauth *OAuthType) CompleteAuthentication(ctx core.Context) error {
	siteInt := ctx.Get("Site")
	site, _ := siteInt.(*OAuthSite)
	sentStateInt := ctx.Get("State")
	method := ctx.Request().Method
	state := ""
	code := ""
	if method == "GET" {
		state = ctx.Query("state")
		code = ctx.Query("code")
	} else {
		req := &OAuthLoginReq{}
		err := ctx.Bind(req)
		if err != nil {
			return err
		}
		state = req.State
		code = req.Code
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "OAuthType: Received code", "state", state, "method", method)
	if state != sentStateInt.(string) {
		log.Logger.Debug(ctx, LOGGING_CONTEXT, "OAuthType: State mismatch", "state", state, "sentStateInt", sentStateInt)
		return errors.ThrowError(ctx, AUTH_ERROR_USER_NOT_FOUND)
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "OAuthType: Received code", "code", code, "method", method)
	oauthCtx := GetOAuthContext(ctx)
	token, err := site.config.Exchange(oauthCtx, code)
	if err != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_INTERNAL_SERVER_ERROR_AUTH, err)
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "OAuthType: Received token", "code", code, "token", token)

	client := site.config.Client(oauthCtx, token)

	log.Logger.Trace(ctx, LOGGING_CONTEXT, "OAuthType: end point profile")

	response, err := client.Get(site.profileURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	//create the user
	usrInt, err := oauth.securityService.CreateUser(ctx)
	if err != nil {
		return errors.RethrowError(ctx, laatoocore.AUTH_ERROR_USEROBJECT_NOT_CREATED, err)
	}

	if err := json.Unmarshal(bits, usrInt); err != nil {
		return errors.RethrowError(ctx, laatoocore.AUTH_ERROR_USEROBJECT_NOT_CREATED, err)
	}

	oauthUsr, ok := usrInt.(auth.OAuthUser)
	if !ok {
		return errors.ThrowError(ctx, laatoocore.AUTH_ERROR_USEROBJECT_NOT_CREATED)
	}

	usrId := fmt.Sprintf("%s_%s", site.sitetype, oauthUsr.GetEmail())

	log.Logger.Debug(ctx, LOGGING_CONTEXT, "OAuthProvider: Authentication Successful", "usrInt", usrInt)

	//get the ide of the user to be tested

	//get the tested user from database
	testedUser, err := oauth.securityService.GetUserById(ctx, usrId)
	if err != nil {
		log.Logger.Info(ctx, LOGGING_CONTEXT, "Tested user not found", "Err", err)
		return errors.RethrowError(ctx, AUTH_ERROR_USER_NOT_FOUND, err)
	}
	if testedUser == nil {
		log.Logger.Info(ctx, LOGGING_CONTEXT, "Tested user not found")
		return errors.ThrowError(ctx, AUTH_ERROR_USER_NOT_FOUND)
	}

	ctx.Set("User", testedUser)
	return nil
}

type OAuthLoginReq struct {
	State string `json:"state" form:"state"`
	Code  string `json:"code" form:"code"`
}
