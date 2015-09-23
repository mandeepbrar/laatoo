package laatooauthentication

import (
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	//"laatoocore"
	"io/ioutil"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/utils"
	"net/http"
	//"net/url"
)

const (
	CONF_AUTHSERVICE_OAUTHPATH_SITES = "oauthsites"
	//login path to be used for local and oauth authentication
	CONF_AUTHSERVICE_OAUTH_AUTHCALLBACK = "callbackurl"
	CONF_AUTHSERVICE_OAUTH_AUTHURL      = "authurl"
	CONF_AUTHSERVICE_OAUTH_TYPE         = "sitetype"
	CONF_AUTHSERVICE_OAUTH_CLIENTID     = "clientid"
	CONF_AUTHSERVICE_OAUTH_CLIENTSECRET = "clientsecret"
)

type OAuthSite struct {
	systemAuthURL string
	callbackURL   string
	config        *oauth2.Config
}

type OAuthType struct {
	sites []*OAuthSite
	//method called in case of callback
	authCallback echo.HandlerFunc
	//reference to the main auth service
	securityService *SecurityService
}

//method called for creating new auth type
func NewOAuth(conf map[string]interface{}, svc *SecurityService) (*OAuthType, error) {
	//create the new auth type
	oauth := &OAuthType{}
	//store the reference to the parent
	oauth.securityService = svc
	log.Logger.Debug(LOGGING_CONTEXT, "OAuthType: Initializing")
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
					return nil, errors.ThrowError(AUTH_ERROR_OAUTH_MISSING_TYPE)
				}
				siteType := siteTypeInt.(string)
				var endpoint oauth2.Endpoint
				switch siteType {
				case "google":
					endpoint = google.Endpoint
				case "facebook":
					endpoint = facebook.Endpoint
				default:
					return nil, errors.ThrowError(AUTH_ERROR_OAUTH_MISSING_TYPE)
				}
				systemAuthUrlInt, ok := siteConf[CONF_AUTHSERVICE_OAUTH_AUTHURL]
				if !ok {
					return nil, errors.ThrowError(AUTH_ERROR_OAUTH_MISSING_AUTHURL)
				}
				clientIdInt, ok := siteConf[CONF_AUTHSERVICE_OAUTH_CLIENTID]
				if !ok {
					return nil, errors.ThrowError(AUTH_ERROR_OAUTH_MISSING_CLIENTID)
				}
				clientSecretInt, ok := siteConf[CONF_AUTHSERVICE_OAUTH_CLIENTSECRET]
				if !ok {
					return nil, errors.ThrowError(AUTH_ERROR_OAUTH_MISSING_CLIENTSECRET)
				}
				callbackURLInt, ok := siteConf[CONF_AUTHSERVICE_OAUTH_AUTHCALLBACK]
				if !ok {
					return nil, errors.ThrowError(AUTH_ERROR_OAUTH_MISSING_CALLBACKURL)
				}
				conf := &oauth2.Config{
					ClientID:     clientIdInt.(string),
					ClientSecret: clientSecretInt.(string),
					RedirectURL:  callbackURLInt.(string),
					Scopes:       []string{"openid", "profile"},
					Endpoint:     endpoint,
				}
				oauth.sites[i] = &OAuthSite{systemAuthURL: systemAuthUrlInt.(string), callbackURL: callbackURLInt.(string), config: conf}
				i++
			}
		}
	}

	return oauth, nil
}

//method called for service
func (oauth *OAuthType) Serve() error {
	return nil
}

//initialize auth type called by base auth for initializing
func (oauth *OAuthType) InitializeType(authStart echo.HandlerFunc, authCallback echo.HandlerFunc) error {
	oauth.authCallback = authCallback
	state := utils.RandomString(10)
	for _, site := range oauth.sites {
		oauth.securityService.Router.Get(site.systemAuthURL, func(ctx *echo.Context) error {
			ctx.Set("Site", site)
			ctx.Set("State", state)
			return authStart(ctx)
		})
		oauth.securityService.Router.Get(site.callbackURL, func(ctx *echo.Context) error {
			ctx.Set("Site", site)
			ctx.Set("State", state)
			return authCallback(ctx)
		})
	}

	return nil
}

//validate the local user
//derive the data from context object
func (oauth *OAuthType) ValidateUser(ctx *echo.Context) error {
	log.Logger.Debug(LOGGING_CONTEXT, "OAuthProvider: Validating Credentials")

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
func (oauth *OAuthType) CompleteAuthentication(ctx *echo.Context) error {

	siteInt := ctx.Get("Site")
	site, _ := siteInt.(*OAuthSite)
	sentStateInt := ctx.Get("State")
	state := ctx.Param("state")
	if state != sentStateInt.(string) {
		return errors.ThrowHttpError(AUTH_ERROR_USER_NOT_FOUND, ctx)
	}
	code := ctx.Param("code")
	token, err := site.config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return err
	}

	client := site.config.Client(oauth2.NoContext, token)

	endpointProfile := ""
	response, err := client.Get(endpointProfile)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	log.Logger.Debug(LOGGING_CONTEXT, "OAuthProvider: Authentication Successful", "bits", bits)
	return nil
}
