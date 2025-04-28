package controllers

import (
	"Project-Golang/models"
	"Project-Golang/repository"
	"Project-Golang/utils"
	"encoding/json"
	"net/http"
	"os"

	"gorm.io/gorm"
)

var Sessions = map[string]string{}

func Register(w http.ResponseWriter, r *http.Request) {
	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse("Invalid request body"))
		return
	}

	if !utils.IsValidEmail(req.Email) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse("Email error"))
		return
	}

	if !utils.IsValidPassword(req.Password) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse("Password error"))
		return
	}

	existingUser, err := repository.GetUserByEmail(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse("Server error"))
		return
	}
	if existingUser != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(models.ErrorResponse("Email already exists"))
		return
	}

	hashed := utils.HashPassword(req.Password)
	user := &models.User{Email: req.Email, Password: hashed}

	if err := repository.CreateUser(user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse("Create account error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.SuccessResponse("Create account success"))
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func Login(w http.ResponseWriter, r *http.Request) {
	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse("Invalid request body"))
		return
	}

	// validate email
	if !utils.IsValidEmail(req.Email) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse("Email error"))
		return
	}

	// get user by email
	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse("Email or password error"))
		return
	}

	// validate password
	hashedInputPassword := utils.HashPassword(req.Password)
	if user.Password != hashedInputPassword {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse("Email or password error"))
		return
	}

	// Create JWT token
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	token, err := utils.GenerateJWT(user.ID, jwtSecret)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse("Authentication error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.SuccessResponse(map[string]string{"token": token}))
}
