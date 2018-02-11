package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
)

// API propeties
type API struct {
	Name        string     `json:"name"`
	Version     *Version   `json:"version"`
	Port        string     `json:"port"`
	Environment string     `json:"environment"`
	Router      chi.Router `json:"-"`
}

// NewAPI constructor
func NewAPI(name string, version string, port string, env string) *API {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	ver, err := NewVersion(version)
	if err != nil {
		log.Fatal(err)
	}

	return &API{
		Name:        name,
		Version:     ver,
		Port:        port,
		Environment: env,
		Router:      router,
	}
}

// Initialize the API
func (api *API) Initialize() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

	api.Router.Get("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bye"))
	})

	//Verify if
	api.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte(" {\"message\":\"Upss!, this route doesn't exist\"} "))
	})

}

// Run the API
func (api *API) Run() {
	socket := ":" + api.Port

	api.Initialize()

	log.Println("Starting API: " + api.ToJSON())
	log.Fatal(http.ListenAndServe(socket, api.Router))
}

// ToJSON return Json string representation of API struct
func (api *API) ToJSON() string {
	json, _ := json.Marshal(api)
	return string(json)
}
