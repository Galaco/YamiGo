package yamigo

import (
	"strings"
	"html/template"
	"net/http"
)

// HTML Template wrapper
type Template struct {
	name string
	templates map[string]string
	parameters map[string]interface{}
}

// Get a parameter added to this template
func (this *Template) GetParameter(key string) interface{} {
	return this.parameters[key]
}

// Render the template
func (this *Template) Render(writer http.ResponseWriter) {
	t,_ := template.New(this.name).Funcs(
		template.FuncMap{
			"Param": this.GetParameter,
		}).ParseFiles(this.templates[this.name],
		this.templates[Configuration.Views.Template.BaseTemplate])

	t.ExecuteTemplate(writer, "base", this.GetParameter)
}

// Add a parameter that template can use
func (this *Template) AddParam(key string, value interface{}) {
	this.parameters[key] = value
}


// Find a matching template by route
func (this *Template) findByName(name string) {
	name = strings.TrimLeft(name, "/")
	this.templates[name] = Configuration.Views.Template.BaseDir + name
}

// Return a new template object
// Attempts to find a template that matches the current path
func NewTemplate(name string) (*Template) {
	t := new(Template)
	t.name = name
	t.templates = make(map[string]string)
	t.parameters = make(map[string]interface{})
	t.findByName(Configuration.Views.Template.BaseTemplate)
	t.findByName(name)

	return t
}