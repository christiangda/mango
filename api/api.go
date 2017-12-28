package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// API propeties
type API struct {
	Name        string     `json:"name"`
	Version     *Version   `json:"version"`
	Port        string     `json:"port"`
	Environment string     `json:"environment"`
	Router      chi.Router `json:"-"`
}

//NewAPI constructor
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

//Initialize the API
func (api *API) Initialize() {

	api.Router.Get("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bye"))
	})

	//Verify if
	api.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte(" {\"message\":\"Upss!, this route doesn't exist\"} "))
	})

}

//Run the API
func (api *API) Run() {
	socket := ":" + api.Port

	log.Println("Initializing API: " + api.ToJSON())
	api.Initialize()

	log.Println("Starting API: " + api.ToJSON())
	log.Fatal(http.ListenAndServe(socket, api.Router))
}

//ToJSON return Json string representation of API struct
func (api *API) ToJSON() string {
	json, _ := json.Marshal(api)
	return string(json)
}
