package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func logging(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		fn(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func baar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "baar")
}

func bee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "bee")
}

// advanced middleware
type Middleware func(http.HandlerFunc) http.HandlerFunc

func createMiddleware(handlerAction http.HandlerFunc) Middleware {
	// Create a new Middleware
	middleware := func(next http.HandlerFunc) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			handlerAction(w, r)
			next(w, r)
		}
		return handler
	}

	// Return newly created middleware
	return middleware
}

func RunMiddlewareServer() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/baar", logging(baar))

	ultraLogger := createMiddleware(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("from ultraLogger ---- ----\n")
		fmt.Printf("%s on %s\n", strings.ToUpper(r.Method), r.URL.Path)
		fmt.Printf("raw query: %v\n", r.URL.RawQuery)
	})

	http.HandleFunc("/bee", ultraLogger(bee))

	port := ":80"
	fmt.Printf("Serving middleware at http://localhost%s/\n", port)
	http.ListenAndServe(port, nil)
}
