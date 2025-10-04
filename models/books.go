package models

import "gorm.io/gorm"

// Book is a struct that represents a book in the database.
type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

// MigrateBooks migrates the Book model to the database.
// It creates the "books" table if it doesn't exist.
func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Book{})
	return err
}
