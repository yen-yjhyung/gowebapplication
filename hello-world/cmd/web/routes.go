package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yen-yjhyung/gowebapplication/hello-world/pkg/config"
	"github.com/yen-yjhyung/gowebapplication/hello-world/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewMux()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
