package auth

const (
	DEFAULT_ROLE = "Role"
)

type Role interface {
	GetId() string
	SetId(string)
	GetIdField() string
	GetPermissions() []string
	SetPermissions([]string)
}
