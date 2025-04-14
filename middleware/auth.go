package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Key để lưu thông tin user vào context
type contextKey string

const UserIDKey contextKey = "user_id"

// AuthMiddleware xác thực JWT token trước khi cho phép truy cập API
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Lấy token từ header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// Kiểm tra định dạng Bearer token
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString := bearerToken[1]
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))

		// Parse và xác thực token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Kiểm tra thuật toán ký
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Lấy thông tin user từ token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Lưu user_id vào context để sử dụng trong handler
		userId, ok := claims["user_id"]
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Tạo context mới với thông tin user
		ctx := context.WithValue(r.Context(), UserIDKey, userId)

		// Gọi đến handler tiếp theo với context đã cập nhật
		next(w, r.WithContext(ctx))
	}
}

// GetUserIDFromContext lấy user ID từ context
func GetUserIDFromContext(ctx context.Context) (float64, bool) {
	userId, ok := ctx.Value(UserIDKey).(float64)
	return userId, ok
}
