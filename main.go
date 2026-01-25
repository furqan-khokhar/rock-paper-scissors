package main

/*
fmt -> i/o
log -> package logging
http ->
*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrader is a struct that will
// take incoming http request
// check if it is a valid ws request
// perform the ws handshake
// return a pointer to a websocket connection (*websocket.Conn)
var upgrader = websocket.Upgrader{
	// data that can be read/written (in bytes per frame)
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// validates http request to be upgraded to ws connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	// infinite loop
	for {
		// tries to read a message from ws conn
		messageType, p, err := conn.ReadMessage()
		// if error
		if err != nil {
			// log error and return
			log.Println(err)
			return
		}

		// log message
		log.Println(string(p))

		// write message back to ws client (echo server)
		if err := conn.WriteMessage(messageType, p); err != nil {
			// if there was an error in writing than print it and return
			log.Println(err)
			return
		}
	}
}

// responds to home request
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

// responds to websocket endpoint request
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrades http to websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	// if there's an error
	if err != nil {
		// log error
		log.Println(err)
	}

	// otherwise connection is successful
	log.Println("Client Connected")

	// send message from server to client
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	// if there's an error
	if err != nil {
		// log error
		log.Println(err)
	}

	// listen indefinitely for new messages
	// through ws connection
	reader(ws)
}

// maps urls to functions
// functions recieve http requests and write http responses
func setupRoutes() {
	// http://localhost:8080/
	http.HandleFunc("/", homePage)
	// http://localhost:8080/ws
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World")
	// sets up routes/endpoints
	setupRoutes()
	// http.ListenAndServe :
	// starts http server on port 8080
	// listen for and serve http requests indefinitely
	// return if there is an error

	// log.Fatal :
	// logs error
	// force exits program
	log.Fatal(http.ListenAndServe(":8080", nil))
}
