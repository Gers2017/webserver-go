package servefiles

import (
	"fmt"
	"net/http"
)

func RunServer() {

	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	port := ":80"
	fmt.Printf("Serving files at http://localhost%s/static/\n", port)
	http.ListenAndServe(port, nil)
}
