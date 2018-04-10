package golaco

import (
	"github.com/julienschmidt/httprouter"
	"strings"
)

type Router struct {
	router *httprouter.Router
	routes []Route
}

func (this *Router) Register(method string, path string, action RouteActionInterface) {
	this.routes = append(this.routes, Route {
		path: path,
		action: action,
	})

	this.router.Handle(strings.ToUpper(method), path, this.routes[len(this.routes)-1].Execute)
}

