package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (MODEL)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Struct for author
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// init books var as a slice Book struct (slice in go is variable length array)
var books []Book

// get all books
func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

// get single book
func getBook(writer http.ResponseWriter, request *http.Request) {

}

// create a new book
func createBook(writer http.ResponseWriter, request *http.Request) {

}

// update book
func updateBook(writer http.ResponseWriter, request *http.Request) {

}

// delete book
func deleteBook(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	// init router
	router := mux.NewRouter()

	// Mock data - @todo - implement db
	books = append(books, Book{ID: "1", Isbn: "438734", Title: "The Hobbit", Author: &Author{Firstname: "JRR", Lastname: "Tolkien"}})
	books = append(books, Book{ID: "2", Isbn: "248734", Title: "The Two Towers", Author: &Author{Firstname: "JRR", Lastname: "Tolkien"}})
	books = append(books, Book{ID: "3", Isbn: "348734", Title: "The Return of the King", Author: &Author{Firstname: "JRR", Lastname: "Tolkien"}})

	// create route handlers/ endpoints, endpoints for API
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
