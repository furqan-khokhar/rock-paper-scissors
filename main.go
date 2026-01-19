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
)

// responds to home request
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

// responds to websocket endpoint request
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
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
	// http.ListenAndServe
	// starts http server on port 8080
	// listen for and serve http requests indefinitely
	// return if there is an error
	// log.Fatal
	// logs error
	// force exits program
	log.Fatal(http.ListenAndServe(":8080", nil))
}
