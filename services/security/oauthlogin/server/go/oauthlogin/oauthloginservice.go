package oauthlogin

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/auth"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	common "securitycommon"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	//"net/http"
)

const (
	//login path to be used for local and oauth authentication
	CONF_SECURITYSERVICE_OAUTH                = "OAUTH"
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
	core.Service
	adminRole       string
	authHeader      string
	tokenGenerator  func(auth.User, string) (string, auth.User, error)
	userDataService data.DataComponent
	userObject      string
	config          *oauth2.Config
	sitetype        string
	systemAuthURL   string
	callbackURL     string
	profileURL      string
	callbackmode    bool
}

/*
func (os *OAuthLoginService) Describe(ctx core.ServerContext) {
	os.SetDescription("Oauth login service")
	os.SetRequestType(config.CONF_OBJECT_STRINGMAP, false, false)
	os.AddStringConfigurations([]string{common.CONF_LOGINSERVICE_USERDATASERVICE}, nil)
	os.AddConfigurations(map[string]string{CONF_OAUTHLOGINSERVICE_SITE: config.CONF_OBJECT_CONFIG})
	os.AddParam(CONF_OAUTHLOGINSERVICE_CALLBACKMODE, config.CONF_OBJECT_BOOL, false)
}*/

func (os *OAuthLoginService) Initialize(ctx core.ServerContext) error {

	sechandler := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sechandler == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	authHeader := sechandler.GetProperty(config.AUTHHEADER)
	if authHeader == nil {
		return errors.ThrowError(ctx, common.AUTH_ERROR_INCORRECT_SECURITY_HANDLER)
	}
	os.authHeader = authHeader.(string)

	os.userObject = sechandler.GetProperty(config.USER).(string)

	return nil

}

//Expects Local user to be provided inside the request
func (os *OAuthLoginService) Invoke(ctx core.RequestContext) error {
	callbackmode, _ := ctx.GetParam(CONF_OAUTHLOGINSERVICE_CALLBACKMODE)
	var res *core.Response
	var err error
	if callbackmode.GetValue().(bool) {
		res, err = os.callbackRequest(ctx)
	} else {
		res, err = os.initialRequest(ctx)
	}
	ctx.SetResponse(res)
	return err
}

//Expects Local user to be provided inside the request
func (os *OAuthLoginService) initialRequest(ctx core.RequestContext) (*core.Response, error) {
	returl, _ := ctx.GetString("callbackurl")
	stateVal, _ := ctx.GetString("state")
	realm, _ := ctx.GetString(config.REALM)

	st := &stateInfo{Url: returl, State: stateVal, Realm: realm}
	state, err := json.Marshal(st)
	if err != nil {
		return core.InternalErrorResponse("Could not marshal to json" + err.Error()), errors.WrapError(ctx, err)
	}
	log.Trace(ctx, "redirecting to url", "state", state)
	encodedState := base64.StdEncoding.EncodeToString(state)
	url := os.config.AuthCodeURL(encodedState)
	log.Trace(ctx, "redirecting to url", "url", url)
	return core.NewServiceResponseWithInfo(core.StatusRedirect, url, nil), nil
}

//Expects Local user to be provided inside the request
func (os *OAuthLoginService) callbackRequest(ctx core.RequestContext) (*core.Response, error) {
	log.Info(ctx, "callback received")
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
	log.Trace(ctx, "received state", "receivedState", st)
	code, ok := ctx.GetString("code")
	if !ok {
		log.Error(ctx, "No code found on oauth return")
		return os.unauthorized(ctx, nil, st.Url)
	}
	log.Trace(ctx, "received code", "code", code)
	return os.authenticate(ctx, code, st)
}

func (os *OAuthLoginService) unauthorized(ctx core.RequestContext, err error, url string) (*core.Response, error) {
	log.Trace(ctx, "unauthorized")
	if url == "" {
		script := []byte(fmt.Sprintf("<html><body onload='var data = {message:\"LoginFailure\"}; window.opener.postMessage(data, \"*\"); window.close();'></body></html>"))
		return core.NewServiceResponseWithInfo(core.StatusServeBytes, &script, map[string]interface{}{core.ContentType: "text/html"}), nil
	} else {
		return core.NewServiceResponseWithInfo(core.StatusRedirect, url, map[string]interface{}{os.authHeader: "", "Error": err}), nil
	}
}

func (os *OAuthLoginService) authenticate(ctx core.RequestContext, code string, st *stateInfo) (*core.Response, error) {
	oauthCtx := ctx.GetOAuthContext()
	token, err := os.config.Exchange(oauthCtx, code)
	if err != nil {
		return os.unauthorized(ctx, errors.WrapError(ctx, err), st.Url)
	}
	log.Trace(ctx, "OAuthType: Received token", "code", code)

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
	usr, err := ctx.CreateObject(os.userObject)
	if err != nil {
		return os.unauthorized(ctx, errors.WrapError(ctx, err), st.Url)
	}

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

	usrs, _, _, recs, err := os.userDataService.Get(ctx, cond, -1, -1, "", nil)
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
		return core.NewServiceResponseWithInfo(core.StatusServeBytes, &script, map[string]interface{}{core.ContentType: "text/html"}), nil
	}

	info := map[string]interface{}{os.authHeader: tokenstr}

	err = ctx.SendSynchronousMessage(common.EVT_LOGIN_SUCCESS, map[string]interface{}{"Data": testedUser, "info": info})
	if err != nil {
		log.Error(ctx, "Encountered Error in sending event", "error", err)
	}

	return core.NewServiceResponseWithInfo(core.StatusRedirect, st.Url, info), nil
}

//Expects Local user to be provided inside the request
func (os *OAuthLoginService) configureSite(ctx core.ServerContext, siteConf config.Config) error {
	siteType, ok := siteConf.GetString(ctx, CONF_OAUTHLOGINSERVICE_OAUTH_TYPE)
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
	clientId, ok := siteConf.GetString(ctx, CONF_OAUTHLOGINSERVICE_CLIENTID)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_OAUTHLOGINSERVICE_CLIENTID)
	}
	clientSecret, ok := siteConf.GetString(ctx, CONF_OAUTHLOGINSERVICE_CLIENTSECRET)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_OAUTHLOGINSERVICE_CLIENTSECRET)
	}
	profile, ok := siteConf.GetString(ctx, CONF_OAUTHLOGINSERVICE_OAUTH_PROFILEURL)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_OAUTHLOGINSERVICE_OAUTH_PROFILEURL)
	}
	callbackURL, ok := siteConf.GetString(ctx, CONF_OAUTHLOGINSERVICE_OAUTH_AUTHCALLBACK)
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

	userDataSvcName, _ := os.GetStringConfiguration(ctx, common.CONF_LOGINSERVICE_USERDATASERVICE)

	userService, err := ctx.GetService(userDataSvcName)
	if err != nil {
		return errors.RethrowError(ctx, common.AUTH_ERROR_MISSING_USER_DATA_SERVICE, err)
	}
	userDataService, ok := userService.(data.DataComponent)
	if !ok {
		return errors.ThrowError(ctx, common.AUTH_ERROR_MISSING_USER_DATA_SERVICE)
	}

	//get and set the data service for accessing users
	os.userDataService = userDataService

	siteconf, _ := os.GetConfiguration(ctx, CONF_OAUTHLOGINSERVICE_SITE)
	return os.configureSite(ctx, siteconf.(config.Config))

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
	log.Trace(ctx, "OAuthType: Received code", "state", state, "method", method)
	if state != sentStateInt.(string) {
		log.Debug(ctx, "OAuthType: State mismatch", "state", state, "sentStateInt", sentStateInt)
		return errors.ThrowError(ctx, AUTH_ERROR_USER_NOT_FOUND)
	}
	log.Trace(ctx, "OAuthType: Received code", "code", code, "method", method)

}


//initialize auth type called by base auth for initializing
func (oauth *OAuthType) InitializeType(ctx core.Context, authStart core.HandlerFunc, authCallback core.HandlerFunc) error {
	oauth.authCallback = authCallback
	for _, site := range oauth.sites {
		log.Debug(ctx, LOGGING_CONTEXT, "OAuthType: Setting up site", "site", site)
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
