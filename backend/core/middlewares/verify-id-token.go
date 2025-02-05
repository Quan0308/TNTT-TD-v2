package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/Quan0308/main-api/interfaces"
	"github.com/Quan0308/main-api/types"
)

func VerifyIdToken(next http.Handler, firebaseAuthService interfaces.FirebaseAuthService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		idToken := parts[1]

		token, err := firebaseAuthService.VerifyIdToken(context.Background(), idToken)
		if err != nil {
			log.Printf("Error verifying ID token: %v\n", err)
			http.Error(w, "Invalid ID token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), types.UserTokenKey, token)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
