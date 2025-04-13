package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"uniqueIndex"`
	Password string
}
