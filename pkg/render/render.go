package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}

func RenderingTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	 err := parsedTemplate.Execute(w, nil)
	 if err != nil{
		fmt.Println("error parsing template", err)
	 }
}

func renderingTest(w http.ResponseWriter) (map[string]*template.Template, error){

	myCache := map[string]*template.Template{}

	pages,err := filepath.Glob("./templates/*.page.html")

	if err != nil{
		fmt.Println("error occured", err)
		return myCache, err
	}

	for _, page := range pages{
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil{
			fmt.Println("error occured", err)
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil{
			fmt.Println("error occured", err)
			return myCache, err
			}

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