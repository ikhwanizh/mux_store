package middleware

import (
	"context"
	"net/http"
	"online-store-backend/config"
	"online-store-backend/helper"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				helper.ResponseError(w, "A token is required for authentication", http.StatusUnauthorized)
				return
			}
			helper.ResponseError(w, "Error reading token", http.StatusBadRequest)
			return
		}

		tokenString := c.Value
		claims := &config.JWTClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		if err != nil {
			helper.ResponseError(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			helper.ResponseError(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
