package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Id       int
	NickName string
}

type UserClaims struct {
	jwt.Claims
	TokenType string
	User
}
