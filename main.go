package main

import (
	"fmt"
	"net/http"
)

// declare portNumber out of main() because I never want the portNumber to be changed by another part of the application.
const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	// http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// start web server that listens for requests
	// listen to particular port of my computer (In this case, listen on 8080)
	_ = http.ListenAndServe(portNumber, nil)
	// access localhost:8080 in browser
}
