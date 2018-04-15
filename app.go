package yamigo

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
)

// YamiGo application
type App struct {
	router Router
}

// Returns router object for YamiGo
func (this *App) Router() *Router{
	return &this.router
}

// Run YamiGo in its current state
func (this *App) Run() {
	host := Configuration.App.Url + ":" + strconv.Itoa(Configuration.App.Port)

	log.Printf("Application running at: %s\n", host)
	log.Fatal(http.ListenAndServe(host, this.Router().router))
}

// Return a new YamiGo applicaton
func New(environment string) *App {
	Configuration.parse(environment)

	app := new(App)
	app.router = Router{
		router: new(httprouter.Router),
	}
	app.router.SetAssetPath(Configuration.Assets.BaseRoute, Configuration.Assets.BaseDir)

	return app
}