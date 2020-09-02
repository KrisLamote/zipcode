package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App ..
type App struct {
	cfg Config
	log *log.Logger
	rtr *mux.Router
}

// NewApp ..
func NewApp(cfg Config, log *log.Logger) *App {
	r := mux.NewRouter()
	r.HandleFunc("/", hello).Methods("GET")

	a := &App{
		cfg: cfg,
		log: log,
		rtr: r,
	}

	return a
}

// Run starts the zipcode http server
func (a *App) Run() {
	a.log.Printf("api : serving on %s\b", a.cfg.API.Host)
	http.ListenAndServe(a.cfg.API.Host, a.rtr)
}

// Shutdown stops the zipcode http server
func (a *App) Shutdown() error {
	a.log.Println("api : shutting down")
	return nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello zipcode")
}
