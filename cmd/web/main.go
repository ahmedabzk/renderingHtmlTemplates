package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ahmed/renderingHtmlTemplates/config"
	"github.com/ahmed/renderingHtmlTemplates/pkg/handlers"
	"github.com/ahmed/renderingHtmlTemplates/pkg/render"
	"github.com/alexedwards/scs/v2"
	
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager




func main(){

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil{
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	fmt.Println("starting server at port", portNumber)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}

}