package main

import (
	"fmt"
	"net/http"

	"github.com/ahmed/renderingHtmlTemplates/pkg/handlers"
)

const portNumber = ":8080"




func main(){
	fmt.Println("starting server at port", portNumber)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(portNumber, nil)

}