package book

import (
	"encoding/json"
	"net/http"
	"time"
)

func AddBook(w http.ResponseWriter, r *http.Request){
	var req Book
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	bookInfo := Book{
		Title: req.Title,
		Author: req.Author,
		Genre: req.Genre,
		Price: req.Price,
		InStock: req.InStock,
		CreatedAt: time.Now(),

	}

	id := InsertBook(bookCollection, bookInfo)
	if id != nil{
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request){
	filter := make(map[string]any)
	books, err := GetBooks(bookCollection, filter)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(books)
	w.WriteHeader(http.StatusOK)
}