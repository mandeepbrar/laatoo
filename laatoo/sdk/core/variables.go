package core

/*server variables default values */
const (
	CONF_SERVER_DEFAULT_ADMINROLE  = "Admin"
	CONF_SERVER_DEFAULT_USEROBJ    = "User"
	CONF_SERVER_DEFAULT_ROLEOBJ    = "Role"
	CONF_SERVER_DEFAULT_AUTHHEADER = "X-Auth-Token"
	CONF_SERVER_DEFAULT_JWTSECRET  = "s;e'c=r@e1t$2D@X2A!"
)

type ServerVariable int

const (
	JWTSECRETKEY ServerVariable = iota
	AUTHHEADER
	ADMINROLE
	USER
	ROLE
)
