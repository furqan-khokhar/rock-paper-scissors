package main

/*
fmt -> i/o
log -> package logging
http -> exchange data between client/server
*/

import (
	"fmt"
	"log"
	"net/http"
)

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
