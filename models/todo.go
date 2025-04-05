package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Todo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"index unique;not null" json:"title" validate:"required,min=10" json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (todo *Todo) Validate() error {
	return validate.Struct(todo)
}

func (todo *Todo) Save(db *gorm.DB) error {
	if err := db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func FindTodoByTitle(db *gorm.DB, title string) (*Todo, error) {
	var todo Todo
	err := db.Where("title = ?", title).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}
