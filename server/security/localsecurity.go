package security

import (
	"crypto/rsa"
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/auth"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	CONF_SECURITY_ROLEDATASERVICE = "role_data_svc"
	CONF_SECURITY_PERMISSIONS     = "permissions"
	CONF_SECURITY_AUTHSERVICES    = "authservices"
	CONF_SECURITY_SUPPORTEDREALMS = "supportedrealms"
)

type realmObj struct {
	name     string
	rolesMap map[string]auth.Role
	//permissions assigned to a role
	rolePermissions map[string]bool
}

type localSecurityHandler struct {
	roleObject      string
	adminRole       string
	anonRole        string
	pvtKey          *rsa.PrivateKey
	roleDataService data.DataComponent
	roleDataSvcName string
	supportedRealms []string
	realms          map[string]*realmObj //all realms
	authServices    []string
	allPermissions  []string
	realmName       string //local realm name
	localRealm      *realmObj
	objectLoader    elements.ObjectLoader
}

func NewLocalSecurityHandler(ctx core.ServerContext, conf config.Config, adminrole string, anonRole string, roleObject string, realmName string) (SecurityPlugin, error) {
	lsh := &localSecurityHandler{adminRole: adminrole, anonRole: anonRole, roleObject: roleObject, realmName: realmName}

	permissions, ok := conf.GetStringArray(ctx, CONF_SECURITY_PERMISSIONS)
	if ok {
		lsh.allPermissions = permissions
	} else {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", CONF_SECURITY_PERMISSIONS)
	}

	supportedRealms, ok := conf.GetStringArray(ctx, CONF_SECURITY_SUPPORTEDREALMS)
	if !ok {
		supportedRealms = []string{realmName}
	}
	lsh.realms = make(map[string]*realmObj, len(supportedRealms))
	log.Info(ctx, "Supported Realms", "realms", supportedRealms)
	for _, supportedRealm := range supportedRealms {
		lsh.realms[supportedRealm] = &realmObj{name: supportedRealm, rolesMap: make(map[string]auth.Role, 15),
			rolePermissions: make(map[string]bool, 100)}
	}

	pvtKeyPath, ok := conf.GetString(ctx, config.CONF_PVTKEYPATH)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", config.CONF_PVTKEYPATH)
	}
	pvtKey, err := utils.LoadPrivateKey(pvtKeyPath)
	if err != nil {
		return nil, errors.RethrowError(ctx, errors.CORE_ERROR_BAD_CONF, err, "conf", config.CONF_PVTKEYPATH)
	}
	lsh.pvtKey = pvtKey

	lsh.objectLoader = ctx.GetServerElement(core.ServerElementLoader).(elements.ObjectLoader)

	roleDataSvcName, ok := conf.GetString(ctx, CONF_SECURITY_ROLEDATASERVICE)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURITY_ROLEDATASERVICE)
	}
	lsh.roleDataSvcName = roleDataSvcName
	authServices, ok := conf.GetStringArray(ctx, CONF_SECURITY_AUTHSERVICES)
	if ok {
		lsh.authServices = authServices
	}
	return lsh, nil
}

func (lsh *localSecurityHandler) Start(ctx core.ServerContext) error {
	roleService, err := ctx.GetService(lsh.roleDataSvcName)
	if err != nil {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURITY_ROLEDATASERVICE)
	}
	roleDataService, ok := roleService.(data.DataComponent)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURITY_ROLEDATASERVICE)
	}
	lsh.roleDataService = roleDataService

	err = lsh.loadRoles(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if lsh.authServices != nil {
		for _, svcName := range lsh.authServices {
			svc, err := ctx.GetService(svcName)
			if err != nil {
				return err
			}
			authComp, ok := svc.(components.AuthenticationComponent)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", CONF_SECURITY_AUTHSERVICES, "svcname", svcName)
			}
			authComp.SetTokenGenerator(ctx, lsh.tokenGenerator(ctx))
		}
	}
	lsh.localRealm = lsh.realms[lsh.realmName]
	log.Info(ctx, "Started local security handler with realms", "realms", lsh.supportedRealms)
	return nil
}

func (lsh *localSecurityHandler) tokenGenerator(ctx core.ServerContext) func(auth.User, string) (string, auth.User, error) {
	return func(user auth.User, realm string) (string, auth.User, error) {
		rbac, ok := user.(auth.RbacUser)
		claims := make(jwt.MapClaims)
		if ok {
			roles, _ := rbac.GetRoles()
			permissions, admin := lsh.GetRolePermissions(roles, realm)
			rbac.SetPermissions(permissions)
			claims["Admin"] = admin
		}
		user.PopulateClaims(claims)
		claims[config.REALM] = realm
		claims["UserId"] = user.GetId()
		//token.Claims["IP"] = ctx.ClientIP()
		claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
		tokenString, err := token.SignedString(lsh.pvtKey)
		if err != nil {
			return "", nil, errors.WrapError(ctx, err)
		}
		return tokenString, user, nil
	}
}

func (lsh *localSecurityHandler) GetRolePermissions(roles []string, realmName string) ([]string, bool) {
	realm, ok := lsh.realms[realmName]
	if !ok {
		return []string{}, false
	}
	permissions := utils.NewStringSet([]string{})
	for _, rolename := range roles {
		role, ok := realm.rolesMap[rolename]
		if ok {
			permissions.Append(role.GetPermissions())
		}
	}
	return permissions.Values(), false
}

func (lsh *localSecurityHandler) HasPermission(ctx core.RequestContext, perm string) bool {
	if lsh.localRealm == nil {
		return false
	}
	return hasPermission(ctx, perm, lsh.localRealm.rolePermissions)
}

func (lsh *localSecurityHandler) AllPermissions(core.RequestContext) []string {
	return lsh.allPermissions
}

func (lsh *localSecurityHandler) loadRoles(ctx core.ServerContext) error {
	adminExists := false
	anonExists := false
	for realmName, realm := range lsh.realms {
		loadRolesReq := ctx.CreateSystemRequest("LoadRoles")
		defer loadRolesReq.CompleteRequest()

		cond, err := lsh.roleDataService.CreateCondition(loadRolesReq, data.FIELDVALUE, map[string]interface{}{"Realm": realmName})
		if err != nil {
			return errors.WrapError(ctx, err)
		}

		roles, _, _, _, err := lsh.roleDataService.Get(loadRolesReq, cond, -1, -1, "", nil)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		for _, val := range roles {
			role, ok := val.(auth.Role)
			if !ok {
				errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
			}
			id := role.GetId()
			roleName := role.GetName()
			realm.rolesMap[roleName] = role
			permissions := role.GetPermissions()
			for _, perm := range permissions {
				key := fmt.Sprintf("%s#%s", roleName, perm)
				realm.rolePermissions[key] = true
			}
			if realmName == lsh.realmName {
				if id == lsh.anonRole {
					anonExists = true
				}
				if id == lsh.adminRole {
					adminExists = true
				}
			}
			log.Trace(ctx, "Loaded role", "rolename", val)
		}
		log.Info(ctx, "Loaded realm", "realm", realmName)
	}
	return lsh.createInitRoles(ctx, anonExists, adminExists, lsh.realmName)
}

func (lsh *localSecurityHandler) createInitRoles(ctx core.ServerContext, anonExists bool, adminExists bool, realm string) error {

	if !anonExists {
		ent, err := lsh.objectLoader.CreateObject(ctx, lsh.roleObject)
		if err != nil {
			return err
		}
		anonymousRole := ent.(auth.Role)
		anonymousRole.SetId(lsh.anonRole)
		anonymousRole.SetRealm(realm)
		anonymousRole.SetName(lsh.anonRole)
		storable, ok := ent.(data.Storable)
		if !ok {
			errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
		}
		anonRolesReq := ctx.CreateSystemRequest("Save Anonymous Role")
		defer anonRolesReq.CompleteRequest()
		err = lsh.roleDataService.Save(anonRolesReq, storable)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	if !adminExists {
		ent, err := lsh.objectLoader.CreateObject(ctx, lsh.roleObject)
		if err != nil {
			return err
		}
		adminRole := ent.(auth.Role)
		adminRole.SetPermissions(lsh.allPermissions)
		adminRole.SetId(lsh.adminRole)
		adminRole.SetRealm(realm)
		adminRole.SetName(lsh.adminRole)
		storable, ok := ent.(data.Storable)
		if !ok {
			errors.ThrowError(ctx, errors.CORE_ERROR_TYPE_MISMATCH)
		}
		adminRolesReq := ctx.CreateSystemRequest("Create Admin Role")
		defer adminRolesReq.CompleteRequest()
		err = lsh.roleDataService.Save(adminRolesReq, storable)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}
