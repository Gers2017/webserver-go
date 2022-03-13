package main

import (
	"fmt"
	"webserver-go/forms"
	"webserver-go/httpserver"
	"webserver-go/json"
	"webserver-go/middleware"
	"webserver-go/mysql"
	"webserver-go/servefiles"
	"webserver-go/sessions"
	"webserver-go/templates"
	"webserver-go/websockets"
)

func main() {
	funcs := []func(){
		httpserver.RunServer,           // 0
		templates.RunTemplate,          // 1
		mysql.RunQueries,               // 2
		servefiles.RunServer,           // 3
		forms.RunServer,                // 4
		websockets.RunServer,           // 5
		middleware.RunMiddlewareServer, // 6
		sessions.RunServer,             // 7
		json.RunServer,                 // 8
	}

	desidedIndex := 8

	for i, f := range funcs {
		if i == desidedIndex {
			fmt.Printf("Running function %d \n", i)
			f()
		}
	}
}
