package main

import (
	"log"
	"net/http"

	"github.com/navdeep0909/bookstore-api/internal/book"
	"github.com/navdeep0909/bookstore-api/internal/user"
)

const port = "8080"

func main(){

	mux := http.NewServeMux()

	mux.HandleFunc("/api/bookstore/signup", user.SignupHandler)
	mux.HandleFunc("/api/bookstore/login", user.LoginHandler)

	//Handler for books
	mux.Handle("/api/bookstore/book", user.JWTAuthMiddleware(user.RequireRole("Admin", http.HandlerFunc(book.AddBookHandler))))
	mux.HandleFunc("/api/bookstore/books", book.GetAllBooksHandler)
	mux.HandleFunc("/api/bookstore/book/{title}", book.GetBookByTitleHandler)
	mux.HandleFunc("/api/bookstore/book/delete/{title}", book.DeleteBookByTitleHandler)
	mux.HandleFunc("/api/bookstore/book/update", book.UpdateBookByTitleHandler)

	// port := os.Getenv("PORT")
    // if port == "" {
    //     port = "8080"
    // }

    log.Printf("🚀 Server listening on http://localhost:%s", port)
    if err := http.ListenAndServe(":"+port, mux); err != nil {
        log.Fatalf("Server failed: %v", err)
    }

}