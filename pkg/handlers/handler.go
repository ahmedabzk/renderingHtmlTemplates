package handlers

import (
	"net/http"

	"github.com/ahmed/renderingHtmlTemplates/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request){
	render.RenderingTemplate(w, "home.page.html")

}

func About(w http.ResponseWriter, r *http.Request){
	render.RenderingTemplate(w, "about.page.html")

}