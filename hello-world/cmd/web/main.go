package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yen-yjhyung/gowebapplication/hello-world/pkg/config"
	"github.com/yen-yjhyung/gowebapplication/hello-world/pkg/handlers"
	"github.com/yen-yjhyung/gowebapplication/hello-world/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCatche()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port", portNumber)
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
