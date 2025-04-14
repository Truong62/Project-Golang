package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT create a JWT token with userID and secret key
func GenerateJWT(userID uint, secret []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID

	expTime := os.Getenv("JWT_EXPIRATION_TIME")
	if expTime == "" {
		expTime = "1h"
	}

	duration, err := time.ParseDuration(expTime)
	if err != nil {
		duration = time.Hour
	}

	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
