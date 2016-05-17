package security

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"laatoo/sdk/auth"
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	//"net/http"
)

const (
	//login path to be used for local and oauth authentication
	CONF_OAUTHLOGINSERVICE_CALLBACKMODE       = "callbackmode"
	CONF_OAUTHLOGINSERVICE_USERDATASERVICE    = "user_data_svc"
	CONF_OAUTHLOGINSERVICE_OAUTH_TYPE         = "sitetype"
	CONF_OAUTHLOGINSERVICE_CLIENTID           = "clientid"
	CONF_OAUTHLOGINSERVICE_CLIENTSECRET       = "clientsecret"
	CONF_OAUTHLOGINSERVICE_KEY                = "key"
	CONF_OAUTHLOGINSERVICE_SITE               = "oauthsite"
	CONF_OAUTHLOGINSERVICE_OAUTH_AUTHURL      = "authurl"
	CONF_OAUTHLOGINSERVICE_OAUTH_AUTHCALLBACK = "callbackurl"
	CONF_OAUTHLOGINSERVICE_OAUTH_PROFILEURL   = "profileurl"
)

type OAuthLoginService struct {
	adminRole       string
	authHeader      string
	tokenGenerator  func(auth.User) (string, auth.User, error)
	userDataService data.DataComponent
	userCreator     core.ObjectCreator
	cipher          cipher.Block
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

	key, ok := conf.GetString(CONF_OAUTHLOGINSERVICE_KEY)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_OAUTHLOGINSERVICE_KEY)
	}
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	os.cipher = cipher

	os.callbackmode, _ = conf.GetBool(CONF_OAUTHLOGINSERVICE_CALLBACKMODE)

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
	if os.callbackmode {
		return os.callbackRequest(ctx)
	} else {
		return os.initialRequest(ctx)
	}
	return nil
}

//Expects Local user to be provided inside the request
func (os *OAuthLoginService) initialRequest(ctx core.RequestContext) error {
	req := ctx.GetRequest()
	body := req.(map[string]interface{})
	returl, ok := body["callbackurl"]
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_REQUEST)
	}
	stateVal, ok := body["state"]
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_REQUEST)
	}
	st := stateInfo{url: returl.(string), state: stateVal.(string)}
	state, err := json.Marshal(st)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	ciphertext := make([]byte, len(state))
	os.cipher.Encrypt(ciphertext, state)
	url := os.config.AuthCodeURL(string(ciphertext))
	ctx.SetResponse(core.NewServiceResponse(core.StatusRedirect, url, nil))
	return nil
}

//Expects Local user to be provided inside the request
func (os *OAuthLoginService) callbackRequest(ctx core.RequestContext) error {
	receivedState, ok := ctx.GetString("State")
	if !ok {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	encBytes := []byte(receivedState)
	jsonBytes := make([]byte, len(encBytes))
	os.cipher.Decrypt(jsonBytes, encBytes)
	st := new(stateInfo)
	err := json.Unmarshal(jsonBytes, st)
	if err != nil {
		errors.WrapError(ctx, err)
		return nil
	}
	code, ok := ctx.GetString("code")
	if !ok {
		log.Logger.Error(ctx, "No code found on oauth return")
		return nil
	}
	return os.authenticate(ctx, code, st)
}

func (os *OAuthLoginService) authenticate(ctx core.RequestContext, code string, st *stateInfo) error {
	oauthCtx := ctx.GetOAuthContext()
	token, err := os.config.Exchange(oauthCtx, code)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Logger.Trace(ctx, "OAuthType: Received token", "code", code, "token", token)

	client := os.config.Client(oauthCtx, token)

	log.Logger.Trace(ctx, "OAuthType: end point profile")

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
		errors.WrapError(ctx, err)
		return nil
	}

	oauthUsr, ok := usr.(auth.OAuthUser)
	if !ok {
		log.Logger.Error(ctx, "Wrong user type")
		return nil
	}

	usrId := fmt.Sprintf("%s_%s", os.sitetype, oauthUsr.GetEmail())

	log.Logger.Debug(ctx, "OAuthProvider: Authorizing user", "usr", usr)

	//get the tested user from database
	testedUser, err := os.userDataService.GetById(ctx, usrId)
	if err != nil || testedUser == nil {
		log.Logger.Info(ctx, "Tested user not found", "Err", err)
		return nil
	}
	tokenstr, _, err := os.tokenGenerator(testedUser.(auth.User))
	if err != nil {
		ctx.SetResponse(core.StatusUnauthorizedResponse)
		return nil
	}
	resp := core.NewServiceResponse(core.StatusRedirect, st.url, map[string]interface{}{os.authHeader: tokenstr})
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
func (os *OAuthLoginService) SetTokenGenerator(ctx core.ServerContext, gen func(auth.User) (string, auth.User, error)) {
	os.tokenGenerator = gen
}
func (os *OAuthLoginService) Start(ctx core.ServerContext) error {
	return nil
}

type stateInfo struct {
	state string
	url   string
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
