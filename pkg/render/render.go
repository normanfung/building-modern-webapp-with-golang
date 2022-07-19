package render

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/normanfung/golang/building-modern-webapp-with-golang/models"
	"github.com/normanfung/golang/building-modern-webapp-with-golang/pkg/config"
)

var app config.AppConfig

func NewTemplates(a config.AppConfig) {
	app = a
}

func Render(w http.ResponseWriter, tmpl string, td models.TemplateData) {
	//Get template cache from app config
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get req template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	//render temp
	err := t.Execute(w, td)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	//range through all files ending with *page.tmpl

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}
	return myCache, nil
}

//Original
// func Render(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles(`./templates/`+tmpl, "templates/base.layout.tmpl")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// }

//Method 1
// var tc = make(map[string]*template.Template)

// func Render(wr http.ResponseWriter, t string) {
// 	//e.g. t -> "home.page.tmpl"
// 	var tmpl *template.Template

// 	//Check if we already have template in our cache
// 	a, inMap := tc[t]

// 	fmt.Println(a, inMap)

// 	if !inMap {
// 		//Need to create the template
// 		log.Println("Creating template adding to cache")
// 		if err := createTemplateCache(t); err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		log.Println("Using cached template")
// 	}

// 	tmpl = tc[t]

// 	if err := tmpl.Execute(wr, nil); err != nil {
// 		log.Println(err)
// 	}

// }

// func createTemplateCache(t string) error {

// 	tmpl, err := template.ParseFiles(`./templates/`+t, "templates/base.layout.tmpl")
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	tc[t] = tmpl
// 	return nil
// }
