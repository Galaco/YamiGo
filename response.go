package yamigo

import "net/http"

// Response object
// Encapsulates a response
type Response struct {
	writer http.ResponseWriter
	template *Template
}

// Set the response template
func (this *Response) SetTemplate(template *Template) {
	this.template = template
}

// Execute the response in its current state
func (this *Response) Execute() {
	this.template.Render(this.writer)
}
