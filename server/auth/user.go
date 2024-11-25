package auth

type TenantInfo interface {
	GetTenantId() string
	GetTenantName() string
}

type User interface {
	GetId() string
	SetId(string)
	GetUsernameField() string
	GetUserName() string
	LoadClaims(map[string]interface{})
	PopulateClaims(map[string]interface{})
	GetEmail() string
	GetRealm() string
	GetTenant() TenantInfo
	GetUserAccount() UserAccount
}
