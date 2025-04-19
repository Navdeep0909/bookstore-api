package main

import (
	"log"
	"net/http"

	"github.com/navdeep0909/bookstore-api/internal/book"
	"github.com/navdeep0909/bookstore-api/internal/user"
)

const port = "8080"

func main(){

	http.HandleFunc("/api/bookstore/signup", user.SignupHandler)
	http.HandleFunc("/api/bookstore/login", user.LoginHandler)

	//Handler for books
	http.HandleFunc("/api/bookstore/book", book.AddBook)
	http.HandleFunc("/api/bookstore/books", book.GetAllBooks)
	http.HandleFunc("/api/bookstore/book/{title}", book.GetBookByTitle)
	http.HandleFunc("/api/bookstore/book/delete/{title}", book.DeleteBookByTitle)

	// port := os.Getenv("PORT")
    // if port == "" {
    //     port = "8080"
    // }

    log.Printf("🚀 Server listening on http://localhost:%s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }

}