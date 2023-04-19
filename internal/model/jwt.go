package model

import "github.com/dgrijalva/jwt-go"

type JWTClaim struct {
	jwt.StandardClaims
	UserID string
}
type ContextData struct {
	UserID string
}
