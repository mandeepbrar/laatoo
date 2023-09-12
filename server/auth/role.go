package auth

type Role interface {
	GetId() string
	SetId(string)
	GetName() string
	SetName(string)
	GetPermissions() []string
	SetPermissions([]string)
	GetTenant() TenantInfo
}
