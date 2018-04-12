package yamigo

import (
	"strings"
	"html/template"
	"net/http"
)

var parameters map[string]interface{}

// Template function to get passed values
func Param(key string) interface{} {
	return parameters[key]
}


// HTML Template wrapper
type Template struct {
	name string
	templates map[string]string
}

func (this *Template) Render(writer http.ResponseWriter) {
	t,_ := template.New(this.name).Funcs(
		template.FuncMap{
			"Param": Param,
		}).ParseFiles(this.templates[this.name],
		this.templates[Configuration.Views.Template.BaseTemplate])

	t.ExecuteTemplate(writer, "base", Param)
}

func (this *Template) AddParam(key string, value interface{}) {
	if parameters == nil {
		parameters = make(map[string]interface{})
	}
	parameters[key] = value
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
	t.findByName(Configuration.Views.Template.BaseTemplate)
	t.findByName(name)

	return t
}