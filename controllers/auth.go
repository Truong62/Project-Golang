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
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !utils.IsValidEmail(req.Email) {
		http.Error(w, "Email error", http.StatusBadRequest)
		return
	}

	if !utils.IsValidPassword(req.Password) {
		http.Error(w, "Password error", http.StatusBadRequest)
		return
	}

	existingUser, err := repository.GetUserByEmail(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		http.Error(w, "Sever error", http.StatusInternalServerError)
		return
	}
	if existingUser != nil {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	hashed := utils.HashPassword(req.Password)
	user := &models.User{Email: req.Email, Password: hashed}

	if err := repository.CreateUser(user); err != nil {
		http.Error(w, "Create account error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.SuccessResponse("Create account success"))
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func Login(w http.ResponseWriter, r *http.Request) {
	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// validate email
	if !utils.IsValidEmail(req.Email) {
		http.Error(w, "Email error", http.StatusBadRequest)
		return
	}

	// get user by email
	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "Email or password error", http.StatusUnauthorized)
		return
	}

	// validate password
	hashedInputPassword := utils.HashPassword(req.Password)
	if user.Password != hashedInputPassword {
		http.Error(w, "Email or password error", http.StatusUnauthorized)
		return
	}

	// Create JWT token
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	token, err := utils.GenerateJWT(user.ID, jwtSecret)
	if err != nil {
		http.Error(w, "Authentication error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.SuccessResponse(map[string]string{"token": token}))
}
