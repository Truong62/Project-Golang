package repository

import (
	"Project-Golang/config"
	"Project-Golang/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}
