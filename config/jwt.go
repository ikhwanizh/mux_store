package config

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var JWT_KEY []byte

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, defaulting to environment variables")
	}
	JWT_KEY = []byte(os.Getenv("JWT_KEY"))
	if len(JWT_KEY) == 0 {
		log.Fatal("JWT_KEY must be set in environment variables")
	}
}

type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
