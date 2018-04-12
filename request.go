package yamigo

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// Request object
// Encapsulates the request
type Request struct {
	request *http.Request
	params httprouter.Params
}

// Get a single parameter from the request
func (this *Request) GetParameter(name string) string {
	return this.params.ByName(name)
}