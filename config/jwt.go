package config

import "github.com/golang-jwt/jwt/v5"

var JWT_Key = []byte("ian123rush456")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
