package yamigo

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

func (this *App) Run() {
	host := Configuration.App.Url + ":" + strconv.Itoa(Configuration.App.Port)
	log.Fatal(http.ListenAndServe(host, this.Router().router))
}


func New(environment string) *App {
	app := new(App)

	app.router = Router{
		router: new(httprouter.Router),
	}
	Configuration.parse(environment)

	return app
}