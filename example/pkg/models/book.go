package models

import (
	"gofr.dev/gofr"

	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func GetAllBooks(c gofr.Context) []Book {
	var books []Book
	c.DB.Find(&books)
	return books
}

func GetBookById(c gofr.Context, id int64) (*Book, error) {
	var book Book
	err := c.DB.Where("ID=?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func CreateBook(c gofr.Context, book *Book) error {
	return c.DB.Create(book).Error
}

func DeleteBook(c gofr.Context, id int64) error {
	return c.DB.Delete(&Book{Model: gorm.Model{ID: id}}).Error
}
