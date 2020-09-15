package internal

import (
	"encoding/json"
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

// structure for a response to a validation or format request
type response struct {
	Country string `json:"country"`
	Format  string `json:"format"`
	Valid   bool   `json:"valid"`
}

// NewApp ..
func NewApp(cfg Config, log *log.Logger) *App {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/validate", validate).Methods("GET")
	r.HandleFunc("/format", validate).Methods("GET")
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

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("route not found"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello zipcode"))
}

func validate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(response{Country: "BE", Format: "3000", Valid: true})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
