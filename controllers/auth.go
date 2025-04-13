package controllers

import (
	"Project-Golang/models"
	"Project-Golang/repository"
	"Project-Golang/utils"
	"net/http"

	"gorm.io/gorm"
)

var Sessions = map[string]string{}

func Register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if !utils.IsValidEmail(email) {
		http.Error(w, "Email error", http.StatusBadRequest)
		return
	}

	if !utils.IsValidPassword(password) {
		http.Error(w, "Password error", http.StatusBadRequest)
		return
	}

	existingUser, err := repository.GetUserByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		http.Error(w, "Sever error", http.StatusInternalServerError)
		return
	}
	if existingUser != nil {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	hashed := utils.HashPassword(password)
	user := &models.User{Email: email, Password: hashed}

	if err := repository.CreateUser(user); err != nil {
		http.Error(w, "Create account error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Create account success"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// validate email
	if !utils.IsValidEmail(email) {
		http.Error(w, "Email error", http.StatusBadRequest)
		return
	}

	// get user by email
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "Email or password error", http.StatusUnauthorized)
		return
	}

	// validate password
	hashedInputPassword := utils.HashPassword(password)
	if user.Password != hashedInputPassword {
		http.Error(w, "Email or password error", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Login success"))
}
