package handlers

import (
	"net/http"

	"github.com/yen-yjhyung/gowebapplication/hello-world/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateV2(w, "home.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateV2(w, "about.html")
}
