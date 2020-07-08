package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books = []Book{
	Book{Id: "1", Title: "Book1", Author: "Author1"},
	Book{Id: "2", Title: "Book2", Author: "Author2"},
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the library!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBooks")
	json.NewEncoder(w).Encode(Books)
}

func returnSingleBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println("Endpoint Hit: returnSingleBook")

	for _, book := range Books {
		if book.Id == id {
			json.NewEncoder(w).Encode(book)
		} else { 
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func createNewBook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)

	fmt.Println("Endpoint Hit: createNewBook")

	Books = append(Books, book)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println("Endpoint Hit: deleteBook")

	for index, book := range Books {
		if book.Id == id {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(Books)
}

func routes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/books", returnAllBooks)
	router.HandleFunc("/book", createNewBook).Methods("POST")
	router.HandleFunc("/book/{id}", returnSingleBook).Methods("GET")
	router.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	return router
}

func handleRequests() {
	router := routes()
	port := ":5000"

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func main() {
	handleRequests()
}
