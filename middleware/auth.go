package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Key to store user information in context
type contextKey string

const UserIDKey contextKey = "user_id"

// AuthMiddleware authenticate JWT token before allowing access to API
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// Check Bearer token format
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString := bearerToken[1]
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))

		// Parse and authenticate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Get user information from token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Save user_id to context to use in handler
		userId, ok := claims["user_id"]
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Create new context with user information
		ctx := context.WithValue(r.Context(), UserIDKey, userId)

		// Call next handler with updated context
		next(w, r.WithContext(ctx))
	}
}

// GetUserIDFromContext get user ID from context
func GetUserIDFromContext(ctx context.Context) (float64, bool) {
	userId, ok := ctx.Value(UserIDKey).(float64)
	return userId, ok
}
