package main

import (
	"fmt"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
)

// declare portNumber out of main() because I never want the portNumber to be changed by another part of the application.
const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	// In develop mode, Usecache sets false because of reloding templates. <- check templates changed.
	// In release mode, Usecashe sets true because of not reloding templates.
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// Use routes
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/divide", Divide)

	// start web server that listens for requests
	// listen to particular port of my computer (In this case, listen on 8080)
	// _ = http.ListenAndServe(portNumber, nil)
	// access localhost:8080 in browser
}

// command in terminal: go run cmd/web/*.go
