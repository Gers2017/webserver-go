package main

//import "webserver-go/httpserver"
// import "webserver-go/templates"
// import "webserver-go/mysql"
// import "webserver-go/servefiles"
// import "webserver-go/forms"
import "webserver-go/websockets"

func main() {
	//httpserver.RunServer()

	// templates.RunTemplate()

	// mysql.RunQueries()

	// servefiles.RunServer()

	// forms.RunServer()

	websockets.RunServer()
}
