package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/arnavsaxena2012/go-restapi/models"
	"github.com/gorilla/mux"
)

var (
	books models.BookStore
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	books = &models.Books{}
	books.Initialize()
}

func main() {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})

	api.HandleFunc("/books/authors/{author}", searchByAuthor).Methods(http.MethodGet)
	api.HandleFunc("/books/book-name/{bookName}", searchByBookName).Methods(http.MethodGet)
	api.HandleFunc("/book/isbn/{isbn}", searchByISBN).Methods(http.MethodGet)
	api.HandleFunc("/book/isbn/{isbn}", deleteByISBN).Methods(http.MethodDelete)
	api.HandleFunc("/book", createBook).Methods(http.MethodPost)
	log.Fatalln(http.ListenAndServe(":8080", r))
}
