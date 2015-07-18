package ginauth_oauth

type User struct {
	RawData           map[string]interface{}
	Email             string
	Name              string
	NickName          string
	Description       string
	UserID            string
	AvatarURL         string
	Location          string
	AccessToken       string
	AccessTokenSecret string
}

type OAuthUser interface {
	GetId() string
	SetId(string)
}
