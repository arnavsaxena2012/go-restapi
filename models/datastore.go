package models

import "github.com/arnavsaxena2012/go-restapi/loader"

type BookStore interface {
	Initialize()
	SearchAuthor(author string, ratingOver, ratingBelow float64, limit, skip int) *[]*loader.BookData
	SearchBook(bookName string, ratingOver, ratingBelow float64, limit, skip int) *[]*loader.BookData
	SearchISBN(isbn string) *loader.BookData
	CreateBook(book *loader.BookData) bool
	DeleteBook(isbn string) bool
	UpdateBook(isbn string, book *loader.BookData) bool
}
