package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Todo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Index     uint      `gorm:"index" json:"index"`
	Title     string    `gorm:"index unique;not null" json:"title" validate:"required,min=10"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (todo *Todo) Validate() error {
	return validate.Struct(todo)
}

func (Todo) TableName() string {
	return "todos"
}

func (todo *Todo) Save(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Exec("LOCK TABLE todos IN EXCLUSIVE MODE").Error; err != nil {
		fmt.Printf("Error locking table: %v\n", err)
		tx.Rollback()
		return err
	}
	fmt.Println("Table locked ... ")

	// Lock table rows for update to avoid race condition
	var maxIndex struct {
		MaxIndex uint
	}
	err := tx.Raw("SELECT COALESCE(MAX(index), 0) as max_index FROM todos").Scan(&maxIndex).Error
	if err != nil {
		fmt.Printf("error: %v\n", err)
		tx.Rollback()
		return err
	}

	// Assign next index
	todo.Index = maxIndex.MaxIndex + 1

	// Save the new todo
	if err := tx.Create(&todo).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
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
