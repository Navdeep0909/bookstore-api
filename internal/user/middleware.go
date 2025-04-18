package user

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("Hustle")

type contextKey string

const (
	UserIDKey contextKey = "userID"
	RoleKey   contextKey = "role"
)

func JWTAuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer"){
			http.Error(w, "Missing or Invalid authorization token", http.StatusUnauthorized)
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
            return JwtKey, nil
        })

		if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
		userID := claims["sub"].(string)
		role := claims["role"].(string)

		// Inject into context
        ctx := context.WithValue(r.Context(), UserIDKey, userID)
		ctx = context.WithValue(ctx, RoleKey, role)
        next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireRole(requiredRole string, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        role, ok := r.Context().Value(RoleKey).(string)
        if !ok || role != requiredRole {
            http.Error(w, "Forbidden: insufficient role", http.StatusForbidden)
            return
        }
        next.ServeHTTP(w, r)
    })
}