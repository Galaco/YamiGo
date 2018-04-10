package golaco

import (
	"strings"
	"html/template"
)


// HTML Template wrapper
type Template struct {
	name string
	template *template.Template
}

// Find a matching template by route
func (this *Template) findByRoute() (*template.Template, error) {
	name := strings.TrimLeft(this.name, "/")
	if name == "" {
		name = "index"
	}
	t,err := template.ParseFiles("tmpl/" + name + ".html")  // Parse template file.

	return t,err
}

// Return internal template object
func (this *Template) Get() *template.Template {
	return this.template
}


// Return a new template object
// Attempts to find a template that matches the current path
func NewTemplate(name string) (*Template,error) {
	ret := new(Template)
	ret.name = name

	t,err := ret.findByRoute()

	if err == nil {
		ret.template = t
	}

	return ret,err
}