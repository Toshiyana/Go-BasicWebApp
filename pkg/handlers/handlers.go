package handlers

import (
	"net/http"

	"myapp/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTamplate(w, "home.page.tmpl")
	// fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTamplate(w, "about.page.tmpl")

	// sum := addValues(2, 2)
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// func addValues(x, y int) int {
// 	return x + y
// }

// // Divide is the divide page handler
// func Divide(w http.ResponseWriter, r *http.Request) {
// 	dividend := 100.0
// 	divisor := 0.0
// 	f, err := divideValues(dividend, divisor)
// 	if err != nil {
// 		fmt.Fprintf(w, "cannot divide by zero")
// 		return
// 	}

// 	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", dividend, divisor, f))
// }

// func divideValues(x, y float64) (float64, error) {
// 	if y == 0 {
// 		err := errors.New("cannot divide by zero")
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }
