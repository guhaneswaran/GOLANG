package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book STRUCT (Model)
type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

//Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init books var as a slice Book struct
var booksarr []Book

//Get all books

func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booksarr)

}

//Get single book
func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get Params

	//Loop thru books and find

	for _, item := range booksarr {
		if item.ID == params["id"] {

			//			fmt.Printf("%d", _a)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})

}

//Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) //Mock ID . not safe as it ll not check for duplicates
	booksarr = append(booksarr, book)
	json.NewEncoder(w).Encode(book)
}

//Update a book - Delete + create combination
func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range booksarr {
		if item.ID == params["id"] {
			booksarr = append(booksarr[:index], booksarr[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			//		book.ID = strconv.Itoa(rand.Intn(1000000)) //Mock ID . not safe as it ll not check for duplicates
			booksarr = append(booksarr, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode(booksarr)

}

//Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range booksarr {
		if item.ID == params["id"] {

			fmt.Println("booksarr[:index] ", booksarr[:index])
			fmt.Println("booksarr[index+1:]", booksarr[index+1:])

			booksarr = append(booksarr[:index], booksarr[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(booksarr)

}

func main() {

	//Init Router
	r := mux.NewRouter()

	//Mock data - @ToDo - implement database
	booksarr = append(booksarr, Book{ID: "1", Isbn: "45679", Title: "Book One",
		Author: Author{Firstname: "John", Lastname: "Doe"}})
	booksarr = append(booksarr, Book{ID: "2", Isbn: "45680", Title: "Book Two",
		Author: Author{Firstname: "Tim", Lastname: "Murphy"}})

	//Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
