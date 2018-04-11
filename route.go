package yamigo

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RouteActionInterface interface {
	Execute(context *RouteContext)
}

type RouteContext struct {
	Writer http.ResponseWriter
	Request *http.Request
	Params httprouter.Params
	Config configuration
}

type Route struct {
	path string
	action RouteActionInterface
}

func (this *Route) RegisterAction(action RouteActionInterface) {
	this.action = action
}

func (this *Route) Execute(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := new(RouteContext)
	ctx.Writer = writer
	ctx.Request = request
	ctx.Params = params
	ctx.Config = Configuration

	this.action.Execute(ctx)
}