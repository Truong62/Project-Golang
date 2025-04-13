package utils

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(pw string) string {
	h := sha256.New()
	h.Write([]byte(pw))
	return fmt.Sprintf("%x", h.Sum(nil))
}
