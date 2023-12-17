package controllers

import (
	"encoding/json"
	"hello/pkg/models"
	"net/http"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	//sending to the postman or the frontend
	w.Write(res)
}
