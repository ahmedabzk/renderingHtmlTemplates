package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request){
	renderingTemplate(w, "home.page.html")

}

func About(w http.ResponseWriter, r *http.Request){
	renderingTemplate(w, "about.page.html")

}

func renderingTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	 err := parsedTemplate.Execute(w, nil)
	 if err != nil{
		fmt.Println("error parsing template", err)
	 }
}

func main(){
	fmt.Println("starting server at port", portNumber)

	http.HandleFunc("/",Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(portNumber, nil)

}