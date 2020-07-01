package main
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
type Book struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
}
var Books []Book
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBooks")
	json.NewEncoder(w).Encode(Books)
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/books", returnAllBooks)
	port := ":5000"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
func main() {
	Books = []Book{
		Book{Title: "Book1", Author: "Author1"},
		Book{Title: "Book2", Author: "Author2"},
	}
	handleRequests()
}