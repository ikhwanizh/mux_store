package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("AES256Key-32Characters1234567890")

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
