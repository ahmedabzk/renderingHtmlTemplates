package handlers

import (
	"net/http"

	"github.com/ahmed/renderingHtmlTemplates/config"
	"github.com/ahmed/renderingHtmlTemplates/models"
	"github.com/ahmed/renderingHtmlTemplates/pkg/render"
)

var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler set the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello there"
	render.RenderingTemplate(w, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderingTemplate(w, "about.page.html", &models.TemplateData{})

}
