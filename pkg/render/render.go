package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ahmed/renderingHtmlTemplates/config"
)

var functions = template.FuncMap{

}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig){
	app = a
}

func RenderingTemplate(w http.ResponseWriter, tmpl string){
	// get the template cache from the app
	tc := app.TemplateCache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("erro")
	}

	buf := new(bytes.Buffer)

	t.Execute(buf, nil)

	_, err := buf.WriteTo(w)

	if err != nil{
		log.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error){
	// cache map that holds all our templates
	myCache := map[string]*template.Template{}

	// get all the path pages that end with .page.html
	pages,err := filepath.Glob("./templates/*.page.html")

	if err != nil{
		fmt.Println("error occured", err)
		return myCache, err
	}

	// loop through the path pages to get the page
	for _, page := range pages{
		// get the name of the page
		name := filepath.Base(page)
		
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil{
			fmt.Println("error occured", err)
			return myCache, err
		}
		// look if there is a match for .layout.html
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil{
			fmt.Println("error occured", err)
			return myCache, err
			}
			// if the len of the matches found is greater than 0
		if len(matches) > 0{
			ts, err := ts.ParseGlob("./templates/*.layout.html")

			if err != nil{
			fmt.Println("error occured", err)
			return myCache, err
			}

			myCache[name] = ts
		}

		
	}

	return myCache, nil
	
}