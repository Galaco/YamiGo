package yamigo

import (
	"strings"
	"html/template"
	"net/http"
)


// HTML Template wrapper
type Template struct {
	name string
	path string
	templates map[string]string
	template *template.Template
}

func (this *Template) Render(writer http.ResponseWriter) {
	this.template = template.Must(
		template.ParseFiles(this.templates[this.name], this.templates[Configuration.Views.Template.BaseTemplate]))

	this.template.ExecuteTemplate(writer, "base", nil)
}



// Find a matching template by route
func (this *Template) findByName(name string) {
	name = strings.TrimLeft(name, "/")
	if name == "" {
		name = "index"
		this.name = "index"
	}
	this.templates[name] = Configuration.Views.Template.BaseDir + name + ".html.tmpl"
}


// Return a new template object
// Attempts to find a template that matches the current path
func NewTemplate(name string) (*Template) {
	t := new(Template)
	t.name = name
	t.templates = make(map[string]string)
	t.findByName(Configuration.Views.Template.BaseTemplate)
	t.findByName(name)

	return t
}