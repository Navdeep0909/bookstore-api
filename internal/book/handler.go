package book

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func AddBookHandler(w http.ResponseWriter, r *http.Request){
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

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request){
	searchString := r.URL.Query().Get("search")
	filter := make(map[string]any)
	if searchString != ""{
		filter["title"] = searchString
	}
	books, err := GetBooks(bookCollection, filter)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(books)
	w.WriteHeader(http.StatusOK)
}

func GetBookByTitleHandler(w http.ResponseWriter, r *http.Request){
	sliceOfUrlStrings := strings.Split(r.URL.Path, "/")
	fmt.Println("Printing the length of slice: ", len(sliceOfUrlStrings))
	title := sliceOfUrlStrings[4]
	fmt.Println("Printing the value in the title var: ", title)
	filter := make(map[string]any)

	filter["title"] = title
	// filter["title"] = {title}
	book, err := GetBookById(bookCollection, &filter)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}

func DeleteBookByTitleHandler(w http.ResponseWriter, r *http.Request){
	sliceOfUrlStrings := strings.Split(r.URL.Path, "/")
	title := sliceOfUrlStrings[5]
	filter := make(map[string]any)

	filter["title"] = title
	// filter["title"] = {title}
	result := DeleteBookById(bookCollection, &filter)
	if result == nil{
		http.Error(w, "Error Occured when trying to delete", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateBookByTitleHandler(w http.ResponseWriter, r *http.Request){
	var req Book
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	fmt.Println("Printing the request object: ", req)
	filter := make(map[string]any)
	filter["title"] = req.Title

	currentBookInfo, err := GetBookById(bookCollection, &filter)
	fmt.Println("Printing the currentBookInfo: ", currentBookInfo)
	if err != nil{
		http.Error(w, "Not found", http.StatusNotFound)
	}
	bookInfo := Book{
		Title: currentBookInfo.Title,
		Author: currentBookInfo.Author,
		Genre: currentBookInfo.Genre,
		Price: currentBookInfo.Price,
		InStock: req.InStock,
		CreatedAt: time.Now(),
	}

	_, err = UpdateBookInfo(bookCollection, filter, bookInfo)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}