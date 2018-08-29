package bootstrap

import (
	"html/template"
	"strings"
)

const defaultLayout = "webapp/views/layouts/default.html"

// GetTemplates load templates with their corresponding layouts and stores
// them in a map; assumes that each template has the definition "Base" in
// at least one of the template files
func GetTemplates() map[string]*template.Template {
	templates := make(map[string]*template.Template)

	templates["home#index"] = loadTemplateWithDefaultLayout("webapp/views/index.html")

	templates["error#error"] = loadTemplate("webapp/views/error.html")

	return templates
}

func loadTemplate(files ...string) *template.Template {

	funcs := template.FuncMap{
		"title": strings.Title,
	}

	return template.Must(
		template.New("").Funcs(funcs).ParseFiles(files...))
}

func loadTemplateWithDefaultLayout(files ...string) *template.Template {

	funcs := template.FuncMap{
		"title": strings.Title,
	}

	allFiles := append([]string{defaultLayout}, files...)

	return template.Must(
		template.New("").Funcs(funcs).ParseFiles(allFiles...))
}
