package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"

	"github.com/gorilla/mux"
)
type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// For the purpose of this exercise we won't worry about persistant data so won't be using a database 
var Books = []Book {
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
        }
	}
}

func createNewBook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	// transform the response body to a new Book object which we can then add to the array aka fake database
	var book Book
	json.Unmarshal(reqBody, &book)
		
	fmt.Println("Endpoint Hit: createNewBook")

    Books = append(Books, book)
    json.NewEncoder(w).Encode(book)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/books", returnAllBooks)
	router.HandleFunc("/book", createNewBook).Methods("POST")
	router.HandleFunc("/book/{id}", returnSingleBook).Methods("GET")
	http.Handle("/",router)
	port := ":5000"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
func main() {
	handleRequests()
}