package yamigo

import (
	"github.com/julienschmidt/httprouter"
	"strings"
	"net/http"
)

// Router wrapper
// Wraps a router implementation into a more convenient struct
type Router struct {
	router *httprouter.Router
	routes []Route
}

// Register a new route
func (this *Router) Register(method string, path string, action RouteActionInterface) {
	this.routes = append(this.routes, Route {
		path: path,
		action: action,
	})

	this.router.Handle(strings.ToUpper(method), path, this.routes[len(this.routes)-1].Execute)
}

func (this *Router) SetAssetPath(route string, filepath string) {
	route = strings.TrimRight(strings.TrimLeft(route, "/"), "/")
	filepath = strings.TrimRight(strings.TrimLeft(filepath, "/"), "/")
	this.router.ServeFiles("/" + route + "/*filepath", http.Dir(filepath))
}

// Route implementation interface
// Any registerable route should fulfill this interface
type RouteActionInterface interface {
	Execute(request Request, response Response)
}

// Embeddable route struct
// Any route struct should embed this
type Route struct {
	path string
	action RouteActionInterface
}

// Register an actionable method to this route
func (this *Route) RegisterAction(action RouteActionInterface) {
	this.action = action
}

// Execute this route
func (this *Route) Execute(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	req := Request {
		request: request,
		params: params,
	}

	res := Response {
		writer: writer,
	}

	this.action.Execute(req, res)
}