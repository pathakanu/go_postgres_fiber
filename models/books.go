package models

import "gorm.io/gorm"

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Book{})
	return err
}
