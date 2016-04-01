package security

/*
import (
	//"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	//"io/ioutil"
	"laatoo/core/registry"
	//"laatoo/sdk/auth"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/data"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	//"net/http"
)

const (
	//login path to be used for local and oauth authentication
	CONF_OAUTHLOGINSERVICE_USERDATASERVICE      = "user_data_svc"
	CONF_OAUTHLOGINSERVICE_OAUTH_TYPE           = "sitetype"
	CONF_OAUTHLOGINSERVICE_CLIENTID             = "clientid"
	CONF_OAUTHLOGINSERVICE_CLIENTSECRET         = "clientsecret"
	CONF_OAUTHLOGINSERVICE_SITE                 = "oauthsite"
	CONF_OAUTHLOGINSERVICE_OAUTH_AUTHURL        = "authurl"
	CONF_OAUTHLOGINSERVICE_OAUTH_AUTHCALLBACK   = "callbackurl"
	CONF_OAUTHLOGINSERVICE_OAUTH_LOGININTERCEPT = "intercept"
	CONF_OAUTHLOGINSERVICE_OAUTH_PROFILEURL     = "profileurl"
)

//service method for doing various tasks
func NewOAuthLoginService(ctx core.ServerContext, conf config.Config) (core.Service, error) {
	return &OAuthLoginService{conf: conf}, nil
}

type OAuthLoginService struct {
	conf        config.Config
	userCreator core.ObjectCreator
	adminRole   string
	//data service to use for users
	UserDataService data.DataService
	config          *oauth2.Config
	sitetype        string
	systemAuthURL   string
	callbackURL     string
	profileURL      string
	state string
	interceptor bool
}

func (os *OAuthLoginService) Initialize(ctx core.ServerContext) error {
	userobject := ctx.GetServerVariable(core.USER)
	userCreator, err := registry.GetObjectCreator(ctx, userobject.(string))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	os.userCreator = userCreator
	userDataSvcName, ok := os.conf.GetString(CONF_OAUTHLOGINSERVICE_USERDATASERVICE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_REGISTRATIONSERVICE_USERDATASERVICE)
	}
	userService, err := ctx.GetService(userDataSvcName)
	if err != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE, err)
	}
	userDataService, ok := userService.(data.DataService)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE)
	}
	log.Logger.Debug(ctx, "User storer set for registration")
	//get and set the data service for accessing users
	os.UserDataService = userDataService
	siteconf, ok := os.conf.GetSubConfig(CONF_OAUTHLOGINSERVICE_SITE)
	os.state := utils.RandomString(10)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_OAUTHLOGINSERVICE_SITE)
	}
	return os.configureSite(ctx, siteconf)
}

//Expects Local user to be provided inside the request
func (os *OAuthLoginService) Invoke(ctx core.RequestContext) error {
	os.initialRequest(ctx)
	os.callbackRequest(ctx)
	return nil
}

//Expects Local user to be provided inside the request
func (os *OAuthLoginService) initialRequest(ctx core.RequestContext) error {
	url := os.config.AuthCodeURL(os.state)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
	return nil
}

//Expects Local user to be provided inside the request
func (os *OAuthLoginService) callbackRequest(ctx core.RequestContext) error {
	receivedState, ok := ctx.GetString("State")
	if !ok || (receivedState != os.state) {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	code, ok := ctx.GetString("code")
	if !ok {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	if (os.interceptor) {
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "OAuthType: Received code", "code", code)
		return ctx.HTML(http.StatusOK, "<html><body onload='var data = {type:\"%s\",state:\"%s\", code:\"%s\"}; window.opener.postMessage(data, \"*\"); window.close();'></body></html", site.sitetype, state, code)
	} else {
		return os.authenticate(ctx)
	}
}

func interceptorRequest() error {
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

}

func (os *OAuthLoginService) authenticate(ctx core.RequestContext, code string) error {

	oauthCtx := core.GetOAuthContext(ctx)
	token, err := os.config.Exchange(oauthCtx, code)
	if err != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_INTERNAL_SERVER_ERROR_AUTH, err)
	}
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "OAuthType: Received token", "code", code, "token", token)

	client := os.config.Client(oauthCtx, token)

	log.Logger.Trace(ctx, LOGGING_CONTEXT, "OAuthType: end point profile")

	response, err := client.Get(os.profileURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	//create the user
	usr, _ := os.userCreator(ctx, nil)

	if err := json.Unmarshal(bits, usr); err != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_USEROBJECT_NOT_CREATED, err)
	}

	oauthUsr, ok := usr.(auth.OAuthUser)
	if !ok {
		return errors.ThrowError(ctx, laatoocore.AUTH_ERROR_USEROBJECT_NOT_CREATED)
	}

	usrId := fmt.Sprintf("%s_%s", os.sitetype, oauthUsr.GetEmail())

	log.Logger.Debug(ctx, "OAuthProvider: Authentication Successful", "usr", usr)

	//get the ide of the user to be tested

	//get the tested user from database
	testedUser, err := os.UserDataService.GetUserById(ctx, usrId)
	if err != nil {
		log.Logger.Info(ctx, LOGGING_CONTEXT, "Tested user not found", "Err", err)
		return errors.RethrowError(ctx, AUTH_ERROR_USER_NOT_FOUND, err)
	}
	if testedUser == nil {
		log.Logger.Info(ctx, LOGGING_CONTEXT, "Tested user not found")
		return errors.ThrowError(ctx, AUTH_ERROR_USER_NOT_FOUND)
	}
	resp, err := completeAuthentication(ctx, existingtestedUserUser)
	if err != nil {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	ctx.SetResponse(resp)
	return nil

}

func (os *OAuthLoginService) GetConf() config.Config {
	return os.conf
}
func (os *OAuthLoginService) GetResponseHandler() core.ServiceResponseHandler {
	return nil
}

//Expects Local user to be provided inside the request
func (os *OAuthLoginService) configureSite(ctx core.ServerContext, siteConf config.Config) error {
	siteType, ok := siteConf.GetString(CONF_OAUTHLOGINSERVICE_OAUTH_TYPE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_OAUTHLOGINSERVICE_OAUTH_TYPE)
	}
	var endpoint oauth2.Endpoint
	switch siteType {
	case "google":
		endpoint = google.Endpoint
	case "facebook":
		endpoint = facebook.Endpoint
	default:
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_OAUTHLOGINSERVICE_OAUTH_TYPE)
	}
	clientId, ok := siteConf.GetString(CONF_OAUTHLOGINSERVICE_CLIENTID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_OAUTHLOGINSERVICE_CLIENTID)
	}
	clientSecret, ok := siteConf.GetString(CONF_OAUTHLOGINSERVICE_CLIENTSECRET)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_OAUTHLOGINSERVICE_CLIENTSECRET)
	}
	profile, ok := siteConf.GetString(CONF_OAUTHLOGINSERVICE_OAUTH_PROFILEURL)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_OAUTHLOGINSERVICE_OAUTH_PROFILEURL)
	}
	callbackURL, ok := siteConf.GetString(CONF_OAUTHLOGINSERVICE_OAUTH_AUTHCALLBACK)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_OAUTHLOGINSERVICE_OAUTH_AUTHCALLBACK)
	}

	conf := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  callbackURL,
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     endpoint,
	}
	interceptor := true
	intercept, ok := siteConf.GetString(CONF_OAUTHLOGINSERVICE_OAUTH_LOGININTERCEPT)
	if ok {
		interceptor = (intercept != "false")
	}
	//oauth.sites[i] = &OAuthSite{sitetype: siteType, interceptor: interceptor, profileURL: profileInt.(string), systemAuthURL: systemAuthUrlInt.(string), callbackURL: callbackURLInt.(string), config: conf}
	os.config = conf
	os.interceptor = interceptor
	os.profileURL = profile
	return nil
}
*/
/*


const (
	CONF_AUTHSERVICE_OAUTHPATH_SITES = "oauthsites"
	//login path to be used for local and oauth authentication
	CONF_AUTHSERVICE_OAUTH_AUTHCALLBACK   = "callbackurl"
	CONF_AUTHSERVICE_OAUTH_LOGINURL       = "oauthlogin"
	CONF_AUTHSERVICE_OAUTH_LOGININTERCEPT = "intercept"
	CONF_AUTHSERVICE_OAUTH_PROFILEURL     = "profileurl"

)



//initialize auth type called by base auth for initializing
func (oauth *OAuthType) InitializeType(ctx core.Context, authStart core.HandlerFunc, authCallback core.HandlerFunc) error {
	oauth.authCallback = authCallback
	for _, site := range oauth.sites {
		log.Logger.Debug(ctx, LOGGING_CONTEXT, "OAuthType: Setting up site", "site", site)
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


//complete authentication
func (oauth *OAuthType) InterceptorPage(ctx core.Context) error {
	siteInt := ctx.Get("Site")
	site, _ := siteInt.(*OAuthSite)
}

//complete authentication
func (oauth *OAuthType) CompleteAuthentication(ctx core.Context) error {
}

*/
/*
type OAuthSite struct {
	interceptor bool
	config      *oauth2.Config
}
type OAuthLoginReq struct {
	State string `json:"state" form:"state"`
	Code  string `json:"code" form:"code"`
}
*/
