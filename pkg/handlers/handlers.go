package handlers

import (
	"net/http"

	"github.com/normanfung/golang/building-modern-webapp-with-golang/models"
	"github.com/normanfung/golang/building-modern-webapp-with-golang/pkg/config"
	"github.com/normanfung/golang/building-modern-webapp-with-golang/pkg/render"
)

//Repo used by handlers
var Repo Repository

// Repository type
type Repository struct {
	App config.AppConfig
}

// NewRepo Create a new repo
func NewRepo(a config.AppConfig) Repository {
	return Repository{
		App: a,
	}
}

//NewHandlers sets the repo for the handlers
func NewHandlers(r Repository) {
	Repo = r
}

func (m Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.Render(w, "home.page.tmpl", models.TemplateData{})
}

func (m Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.Render(w, "about.page.tmpl", models.TemplateData{StringMap: stringMap})
}
