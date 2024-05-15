package models

import "github.com/dgrijalva/jwt-go"

type AccessJWTClaims struct {
	jwt.StandardClaims
	UserID int
}

type RefreshJWTClaims struct {
	jwt.StandardClaims
	UserID int
}

type JWTPairTokens struct {
	AccessToken  string
	RefreshToken string
}
