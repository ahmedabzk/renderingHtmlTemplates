package config

import "html/template"

// AppConfig holds application configuration
type AppConfig struct{
	TemplateCache map[string]*template.Template
}