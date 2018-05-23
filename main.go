package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Price  *Price  `json:"price"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Biography string `json:"biography"`
}

// Price Struct
type Price struct {
	US float64 `json:"us-price"`
	UK float64 `json:"uk-price"`
	AU float64 `json:"au-price"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
// All Handle functions needs to have a writer and request inputs
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get a Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// loop through books and find id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create Books
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not good for Prod
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
}
func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data
	// TODO - impliment DB
	books = append(books, Book{ID: "1", Isbn: "448743", Price: &Price{US: 9.99, UK: 7.83, AU: 11.88}, Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe", Biography: "Likes to write horror and take long walks on sandy beaches."}})
	books = append(books, Book{ID: "2", Isbn: "847564", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith", Biography: "Likes to write romance and go cliff diving."}})
	books = append(books, Book{ID: "3", Isbn: "123564", Title: "Book Three", Author: &Author{Firstname: "Lara", Lastname: "Johnson", Biography: "Likes to write poetry and is in a metal band."}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
