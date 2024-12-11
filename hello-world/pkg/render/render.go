package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
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

func RenderTemplateV2(w http.ResponseWriter, html string) {
	// 1. create a template cache
	tc, err := createTemplateCatche()
	if err != nil {
		log.Fatal(err)
	}

	// 2. get requested template from cache
	t, ok := tc[html]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// 3. render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCatche() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from ./hello-world/templates
	pages, err := filepath.Glob("./hello-world/templates/*.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./hello-world/templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./hello-world/templates/*.layout.html")
			if err != nil {
				return myCache, nil
			}
		}

		myCache[name] = ts

	}

	return myCache, nil

}
