package controllers

import (
	"net/http"

	"gofr.dev/gofr"

	"hello/pkg/models"
)

var NewBook models.Book

func GetBooks(c gofr.Context) []models.Book {
	return models.GetAllBooks(c.DB)
}

func GetBookByID(c gofr.Context, id int64) (*models.Book, error) {
	bookDetails, err := models.GetBookByID(c.DB, id)
	if err != nil {
		return nil, err
	}

	return bookDetails, nil
}

func CreateBook(c gofr.Context) (*models.Book, error) {
	book := &models.Book{}
	if err := c.BodyParser(book); err != nil {
		return nil, err
	}

	createdBook, err := book.CreateBook(c.DB)
	if err != nil {
		return nil, err
	}
	return createdBook, nil
}

func DeleteBookByID(c gofr.Context, id int64) error {
	err := models.DeleteBookByID(c.DB, id)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, map[string]string{"message": "Book deleted successfully"})
	return nil
}

// update function
func UpdateBookByID(c gofr.Context, id int64) (*models.Book, error) {
	updateBook := &models.Book{}
	if err := c.BodyParser(updateBook); err != nil {
		return nil, err
	}

	bookDetails, err := models.GetBookByID(c.DB, id)
	if err != nil {
		return nil, err
	}

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	if err := c.DB.Save(&bookDetails); err != nil {
		return nil, err
	}

	return bookDetails, nil
}
