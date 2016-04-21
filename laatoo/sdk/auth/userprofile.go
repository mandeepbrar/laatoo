package auth

type UserProfile interface {
	GetId() string
	SetId(string)
	GetIdField() string
	GetEmail() string
	GetName() string
	GetPicture() string
	GetGender() string
}