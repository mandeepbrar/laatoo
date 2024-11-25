package auth

type UserAccount interface {
	GetId() string
	GetEmail() string
	GetFullName() string
	GetPicture() string
	GetGender() string
	GetUsernameField() string
	GetUserName() string
}
