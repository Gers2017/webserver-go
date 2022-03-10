package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"] // the book title slug
	page := vars["page"]   // the page
	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	r.HandleFunc("/books/{title}/page/{page}", HandleBooks).Methods("GET")

	fmt.Println("Listening to port 5050")
	http.ListenAndServe(":5050", r)
}
