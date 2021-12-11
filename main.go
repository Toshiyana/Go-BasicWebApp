package main

import (
	"fmt"
	"net/http"
)

func main() {
	// r: request type use pointer
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Don't use Println() because it print in console
		n, err := fmt.Fprintf(w, "Hello world")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Number of bytes written: %d\n", n)
	})

	// start web server that listens for requests
	// listen to particular port of my computer (In this case, listen on 8080)
	_ = http.ListenAndServe(":8080", nil)
	// access localhost:8080 in browser
}
