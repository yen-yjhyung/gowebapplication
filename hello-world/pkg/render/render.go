package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./hello-world/templates/"+html, "./hello-world/templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

var templatecache = make(map[string]*template.Template)

func RenderTemplateV2(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := templatecache[t]
	if !inMap {
		// need to create new template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}

	tmpl = templatecache[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./hello-world/templates/%s", t),
		"./hello-world/templates/base.layout.html",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	templatecache[t] = tmpl

	return nil
}
