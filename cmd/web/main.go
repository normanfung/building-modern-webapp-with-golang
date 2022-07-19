package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/normanfung/golang/building-modern-webapp-with-golang/pkg/config"
	"github.com/normanfung/golang/building-modern-webapp-with-golang/pkg/handlers"
	"github.com/normanfung/golang/building-modern-webapp-with-golang/pkg/render"
)

const portNumber = "localhost:8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	//SessionManager
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.UseCache = false
	app.TemplateCache = tc

	//Set app.UseCache for render
	render.NewTemplates(app)
	//Return a repository type
	repo := handlers.NewRepo(app)
	handlers.NewHandlers(repo)

	http.ListenAndServe(portNumber, session.LoadAndSave(routes(app)))
}
