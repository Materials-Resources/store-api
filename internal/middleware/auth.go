package middleware

import (
	"customer-api/internal/session"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type Claims struct {
	Metadata map[string]string `json:"urn:zitadel:iam:user:metadata"`
	jwt.RegisteredClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get Authorization header
		bearer := r.Header.Get("Authorization")

		// Remove "Bearer " prefix
		tokenString := strings.TrimPrefix(bearer, "Bearer: ")

		_, err := session.ParseJwt(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
