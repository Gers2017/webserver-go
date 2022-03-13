package websockets

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reverseMessage(message string) string {
	runes := []rune(message)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func RunServer() {

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil {
			log.Fatal(err)
		}

		for {
			// Read message from browser
			msgType, msg, _ := ws.ReadMessage()

			// Print the message to the console
			fmt.Printf("%s says: %s\n", ws.RemoteAddr(), string(msg))
			revMsg := reverseMessage(string(msg))
			// Write message back to browser
			if err = ws.WriteMessage(msgType, []byte(revMsg)); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "text/html")
		http.ServeFile(response, request, "./public/websockets.html")
	})

	//fs := http.FileServer(http.Dir("public/"))
	//http.Handle("/", http.StripPrefix("/", fs))

	port := ":8080"

	fmt.Printf("Websockets, Serving files at http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
