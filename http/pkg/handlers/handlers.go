package handlers

import (
	"net/http"

	"github.com/Nguyenthaiduc/go-web/pkg/render"
)

// Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
