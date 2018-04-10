package golaco

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
)

type App struct {
	router Router
}

func (this *App) Router() *Router{
	return &this.router
}

func (app *App) Run() {
	host := Configuration.App.Url + ":" + strconv.Itoa(Configuration.App.Port)
	log.Fatal(http.ListenAndServe(host, app.Router().router))
}


func Golaco(environment string) *App {
	app := new(App)

	app.router = Router{
		router: new(httprouter.Router),
	}
	Configuration.parse(environment)

	return app
}