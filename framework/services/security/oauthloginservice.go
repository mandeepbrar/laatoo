package security

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"laatoo/sdk/auth"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	//"net/http"
)

const (
	//login path to be used for local and oauth authentication
	CONF_OAUTHLOGINSERVICE_CALLBACKMODE       = "callbackmode"
	CONF_OAUTHLOGINSERVICE_USERDATASERVICE    = "user_data_svc"
	CONF_OAUTHLOGINSERVICE_OAUTH_TYPE         = "sitetype"
	CONF_OAUTHLOGINSERVICE_CLIENTID           = "clientid"
	CONF_OAUTHLOGINSERVICE_CLIENTSECRET       = "clientsecret"
	CONF_OAUTHLOGINSERVICE_SITE               = "oauthsite"
	CONF_OAUTHLOGINSERVICE_OAUTH_AUTHURL      = "authurl"
	CONF_OAUTHLOGINSERVICE_OAUTH_AUTHCALLBACK = "callbackurl"
	CONF_OAUTHLOGINSERVICE_OAUTH_PROFILEURL   = "profileurl"
)

type OAuthLoginService struct {
	adminRole       string
	authHeader      string
	tokenGenerator  func(auth.User, string) (string, auth.User, error)
	userDataService data.DataComponent
	userCreator     core.ObjectCreator
	config          *oauth2.Config
	sitetype        string
	systemAuthURL   string
	callbackURL     string
	profileURL      string
	callbackmode    bool
}

func (os *OAuthLoginService) Initialize(ctx core.ServerContext, conf config.Config) error {
	sechandler := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sechandler == nil {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	userDataSvcName, ok := conf.GetString(CONF_LOGINSERVICE_USERDATASERVICE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_LOGINSERVICE_USERDATASERVICE)
	}
	userService, err := ctx.GetService(userDataSvcName)
	if err != nil {
		return errors.RethrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE, err)
	}
	userDataService, ok := userService.(data.DataComponent)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_MISSING_USER_DATA_SERVICE)
	}
	authHeader, ok := sechandler.GetString(config.AUTHHEADER)
	if !ok {
		return errors.ThrowError(ctx, AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	os.authHeader = authHeader

	userObject, ok := sechandler.GetString(config.USER)
	if !ok {
		userObject = config.DEFAULT_USER
	}

	userCreator, err := ctx.GetObjectCreator(userObject)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	os.userCreator = userCreator

	//get and set the data service for accessing users
	os.userDataService = userDataService
	siteconf, ok := conf.GetSubConfig(CONF_OAUTHLOGINSERVICE_SITE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_OAUTHLOGINSERVICE_SITE)
	}
	return os.configureSite(ctx, siteconf)
}

//Expects Local user to be provided inside the request
func (os *OAuthLoginService) Invoke(ctx core.RequestContext) error {
	callbackmode, _ := ctx.GetBool(CONF_OAUTHLOGINSERVICE_CALLBACKMODE)
	if callbackmode {
		return os.callbackRequest(ctx)
	} else {
		return os.initialRequest(ctx)
	}
}

//Expects Local user to be provided inside the request
func (os *OAuthLoginService) initialRequest(ctx core.RequestContext) error {
	returl, _ := ctx.GetString("callbackurl")
	stateVal, _ := ctx.GetString("state")
	realm, _ := ctx.GetString(config.REALM)

	st := &stateInfo{Url: returl, State: stateVal, Realm: realm}
	state, err := json.Marshal(st)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Logger.Trace(ctx, "redirecting to url", "state", state)
	encodedState := base64.StdEncoding.EncodeToString(state)
	url := os.config.AuthCodeURL(encodedState)
	log.Logger.Trace(ctx, "redirecting to url", "url", url)
	ctx.SetResponse(core.NewServiceResponse(core.StatusRedirect, url, nil))
	return nil
}

//Expects Local user to be provided inside the request
func (os *OAuthLoginService) callbackRequest(ctx core.RequestContext) error {
	log.Logger.Info(ctx, "callback received")
	receivedState, ok := ctx.GetString("state")
	if !ok {
		return os.unauthorized(ctx, nil, "")
	}
	decodedState, err := base64.StdEncoding.DecodeString(receivedState)
	st := new(stateInfo)
	err = json.Unmarshal(decodedState, st)
	if err != nil {
		return os.unauthorized(ctx, errors.WrapError(ctx, err), "")
	}
	log.Logger.Trace(ctx, "received state", "receivedState", st)
	code, ok := ctx.GetString("code")
	if !ok {
		log.Logger.Error(ctx, "No code found on oauth return")
		return os.unauthorized(ctx, nil, st.Url)
	}
	log.Logger.Trace(ctx, "received code", "code", code)
	return os.authenticate(ctx, code, st)
}

func (os *OAuthLoginService) unauthorized(ctx core.RequestContext, err error, url string) error {
	log.Logger.Trace(ctx, "unauthorized")
	ctx.SetResponse(core.StatusUnauthorizedResponse)
	if url == "" {
		script := []byte(fmt.Sprintf("<html><body onload='var data = {message:\"LoginFailure\"}; window.opener.postMessage(data, \"*\"); window.close();'></body></html>"))
		resp := core.NewServiceResponse(core.StatusServeBytes, &script, map[string]interface{}{core.ContentType: "text/html"})
		ctx.SetResponse(resp)

	} else {
		resp := core.NewServiceResponse(core.StatusRedirect, url, map[string]interface{}{os.authHeader: "", "Error": err})
		ctx.SetResponse(resp)
	}
	return nil
}

func (os *OAuthLoginService) authenticate(ctx core.RequestContext, code string, st *stateInfo) error {
	oauthCtx := ctx.GetOAuthContext()
	token, err := os.config.Exchange(oauthCtx, code)
	if err != nil {
		return os.unauthorized(ctx, errors.WrapError(ctx, err), st.Url)
	}
	log.Logger.Trace(ctx, "OAuthType: Received token", "code", code)

	client := os.config.Client(oauthCtx, token)

	response, err := client.Get(os.profileURL)
	if err != nil {
		return os.unauthorized(ctx, errors.WrapError(ctx, err), st.Url)
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return os.unauthorized(ctx, errors.WrapError(ctx, err), st.Url)
	}

	//create the user
	usr := os.userCreator()

	if err := json.Unmarshal(bits, usr); err != nil {
		return os.unauthorized(ctx, errors.WrapError(ctx, err), st.Url)
	}

	oauthUsr, ok := usr.(auth.OAuthUser)
	if !ok {
		return os.unauthorized(ctx, errors.WrapError(ctx, err), st.Url)
	}

	usrId := fmt.Sprintf("%s_%s", os.sitetype, oauthUsr.GetEmail())

	argsMap := map[string]interface{}{oauthUsr.GetUsernameField(): usrId, config.REALM: st.Realm}

	cond, err := os.userDataService.CreateCondition(ctx, data.FIELDVALUE, argsMap)
	if err != nil {
		return os.unauthorized(ctx, errors.WrapError(ctx, err), st.Url)
	}

	usrs, _, _, recs, err := os.userDataService.Get(ctx, cond, -1, -1, "", "")
	if err != nil || recs <= 0 {
		return os.unauthorized(ctx, err, st.Url)
	}

	//get the tested user from database
	testedUser := usrs[0].(auth.User)

	tokenstr, _, err := os.tokenGenerator(testedUser, st.Realm)
	if err != nil {
		return os.unauthorized(ctx, errors.WrapError(ctx, err), st.Url)
	}
	if st.Url == "" {
		permissions, _ := testedUser.(auth.RbacUser).GetPermissions()
		permissionsArr, _ := json.Marshal(permissions)
		script := []byte(fmt.Sprintf("<html><body onload='var data = {message:\"LoginSuccess\", token:\"%s\", id:\"%s\", permissions:%s}; window.opener.postMessage(data, \"*\"); window.close();'></body></html>", tokenstr, testedUser.GetId(), string(permissionsArr)))
		resp := core.NewServiceResponse(core.StatusServeBytes, &script, map[string]interface{}{core.ContentType: "text/html"})
		ctx.SetResponse(resp)
		return nil
	}
	resp := core.NewServiceResponse(core.StatusRedirect, st.Url, map[string]interface{}{os.authHeader: tokenstr})
	ctx.SetResponse(resp)
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
	os.sitetype = siteType
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
	//oauth.sites[i] = &OAuthSite{sitetype: siteType, interceptor: interceptor, profileURL: profileInt.(string), systemAuthURL: systemAuthUrlInt.(string), callbackURL: callbackURLInt.(string), config: conf}
	os.config = conf
	os.profileURL = profile
	return nil
}
func (os *OAuthLoginService) SetTokenGenerator(ctx core.ServerContext, gen func(auth.User, string) (string, auth.User, error)) {
	os.tokenGenerator = gen
}
func (os *OAuthLoginService) Start(ctx core.ServerContext) error {
	return nil
}

type stateInfo struct {
	State string
	Url   string
	Realm string
}

/*


const (
	CONF_AUTHSERVICE_OAUTHPATH_SITES = "oauthsites"
	//login path to be used for local and oauth authentication
	CONF_AUTHSERVICE_OAUTH_AUTHCALLBACK   = "callbackurl"
	CONF_AUTHSERVICE_OAUTH_LOGINURL       = "oauthlogin"
	CONF_AUTHSERVICE_OAUTH_LOGININTERCEPT = "intercept"
	CONF_AUTHSERVICE_OAUTH_PROFILEURL     = "profileurl"

)

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
	log.Logger.Trace(ctx, "OAuthType: Received code", "state", state, "method", method)
	if state != sentStateInt.(string) {
		log.Logger.Debug(ctx, "OAuthType: State mismatch", "state", state, "sentStateInt", sentStateInt)
		return errors.ThrowError(ctx, AUTH_ERROR_USER_NOT_FOUND)
	}
	log.Logger.Trace(ctx, "OAuthType: Received code", "code", code, "method", method)

}


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