package main

import (
	"errors"
	"fmt"
	"net/http"
)

// declare portNumber out of main() because I never want the portNumber to be changed by another part of the application.
const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

// Divide is the divide page handler
func Divide(w http.ResponseWriter, r *http.Request) {
	dividend := 100.0
	divisor := 0.0
	f, err := divideValues(dividend, divisor)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", dividend, divisor, f))
}

func divideValues(x, y float64) (float64, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// start web server that listens for requests
	// listen to particular port of my computer (In this case, listen on 8080)
	_ = http.ListenAndServe(portNumber, nil)
	// access localhost:8080 in browser
}
