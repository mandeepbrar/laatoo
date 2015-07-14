package ginauth_local

type LocalAuthUser interface {
	GetId() string
	SetId(string)
	GetPassword() string
}
