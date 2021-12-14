package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// Set base cookie because it uses cookies to make sure that the token it generates is available on a per page basis.
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",   // apply to the entire site for cookie secure
		Secure:   false, // In production, change true
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
