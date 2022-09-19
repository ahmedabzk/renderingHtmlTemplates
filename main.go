package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request){

}

func About(w http.ResponseWriter, r *http.Request){

}

func main(){
	fmt.Println("starting server at port", portNumber)

	http.HandleFunc("/",Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(portNumber, nil)

}