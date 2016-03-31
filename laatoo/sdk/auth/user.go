package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	DEFAULT_USER = "User"
)

type User interface {
	GetId() string
	SetId(string)
	GetIdField() string
	SetJWTClaims(*jwt.Token)
	LoadJWTClaims(*jwt.Token)
}
