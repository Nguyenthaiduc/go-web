package handlers

import (
	"net/http"

	"github.com/Nguyenthaiduc/go-web/pkg/render"
	"github.com/Nguyenthaiduc/go-web/pkg/config"
	
)
//Repo the repository used by handlers
var Repo *Repository
//type Repository type
type Repository struct {
	App *config.AppConfig

}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App : a,
	}
}
//NewHandler Repository sets the repository the handlers
func NewHandler(r *Repository) {
	Repo = r
}


// Home is the about page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
