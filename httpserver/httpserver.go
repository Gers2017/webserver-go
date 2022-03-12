package httpserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var books map[string]int = map[string]int{
	"java-book":  2000,
	"c-language": 12300,
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "all the books: %v\n", books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	pages := books[title]
	fmt.Fprintf(w, "Book pages: %d", pages)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	// q := r.URL.Query()
	// title := q.Get("title")
	// pages := q.Get("pages")
	vars := mux.Vars(r)
	title := vars["title"]
	pages := vars["pages"]

	value, err := strconv.Atoi(pages)
	if err != nil {
		value = 0
	}
	books[title] = int(value)
	fmt.Fprintf(w, "You've created the book: %s with %d pages\n", title, value)
}

func RunServer() {
	fs := http.FileServer(http.Dir("httpserver/static/"))     // actual path of your files
	http.Handle("/static/", http.StripPrefix("/static/", fs)) //server endpoint

	// gorilla mux httpserver
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/books/{title}", GetBook).Methods("GET")
	r.HandleFunc("/books/{title}/pages/{pages}", AddBook).Methods("GET")
	r.HandleFunc("/books/", AllBooks).Methods("GET")

	port := ":80"
	fmt.Printf("Listening at http://localhost%s/static/\n", port)
	http.ListenAndServe(port, nil)
}
