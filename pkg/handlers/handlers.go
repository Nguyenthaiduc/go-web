package handlers

import (
	"net/http"

	"github.com/Nguyenthaiduc/go-web/pkg/config"
	"github.com/Nguyenthaiduc/go-web/pkg/render"
	"github.com/Nguyenthaiduc/go-web/pkg/models"
)

//templateData holds


//Repo the repository used by handlers
var Repo *Repository

//type Repository type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandler Repository sets the repository the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// Home is the about page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIP) 

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello,again"

	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")

	

	stringMap["remote_ip"] = remoteIP


	//send

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
