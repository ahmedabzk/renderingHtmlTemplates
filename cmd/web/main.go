package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ahmed/renderingHtmlTemplates/config"
	"github.com/ahmed/renderingHtmlTemplates/pkg/handlers"
	"github.com/ahmed/renderingHtmlTemplates/pkg/render"
)

const portNumber = ":8080"




func main(){
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil{
		log.Fatal(err)
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	fmt.Println("starting server at port", portNumber)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(portNumber, nil)

}