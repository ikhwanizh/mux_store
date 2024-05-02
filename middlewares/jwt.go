package middlewares

import (
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
				helper.ResponseError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}
		// mengambil token value
		tokenString := c.Value

		claims := &config.JWTClaims{}
		// parsing token jwt
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			helper.ResponseError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			helper.ResponseError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
