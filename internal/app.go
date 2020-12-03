package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/KrisLamote/zipcode/engine"
	"github.com/KrisLamote/zipcode/yaml"
	"github.com/gorilla/mux"
)

// App ..
type App struct {
	cfg Config
	eng *engine.Engine
	log *log.Logger
	rtr *mux.Router
}

// structure for a request for validation or format
type request struct {
	Country string `json:"country"`
	Zipcode string `json:"zipcode"`
}

// structure for a response to a validation or format request
type response struct {
	Country string `json:"country"`
	Format  string `json:"format"`
	Valid   bool   `json:"valid"`
}

// NewApp ..
func NewApp(cfg Config, log *log.Logger) *App {
	data, err := ioutil.ReadFile("../../data/rules.yml")
	if err != nil {
		log.Printf("api : config failed reading rules from yml, error: %s\n", err)
		data = []byte("BE: [\"####\"]\nBR: [\"#####-###\", \"#####\"]\nSK: [\"## ###\"]")
	}

	rules, err := yaml.Parse(data)
	if err != nil {
		log.Printf("api : config failed reading rules any rules, error: %s\n", err)
		os.Exit(1)
	}

	a := &App{
		cfg: cfg,
		eng: engine.New(rules),
		log: log,
	}

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/validate", a.validate).Methods("POST")
	r.HandleFunc("/format", a.validate).Methods("POST")
	r.HandleFunc("/", hello).Methods("GET")
	a.rtr = r

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

func (a *App) validate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if we have both country & zipcide
	if len(req.Country) == 0 || len(req.Zipcode) == 0 {
		http.Error(w, "either country or zip are missing", http.StatusBadRequest)
		return
	}

	valid := a.eng.Valid(req.Zipcode, req.Country)
	body, err := json.Marshal(response{Country: req.Country, Format: req.Zipcode, Valid: valid})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
