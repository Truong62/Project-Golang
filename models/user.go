package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"uniqueIndex"`
	Password string
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
