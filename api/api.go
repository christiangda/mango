package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// API propeties
type API struct {
	Name    string     `json:"name"`
	Version *Version   `json:"version"`
	Port    string     `json:"port"`
	Router  chi.Router `json:"-"`
}

//NewAPI constructor
func NewAPI(name string, version string, port string) (*API, error) {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	ver, err := NewVersion(version)
	if err != nil {
		return nil, errors.New("api: bad string version format")
	}

	return &API{
		Name:    name,
		Version: ver,
		Port:    port,
		Router:  router,
	}, nil
}

//Initialize the API
func (api *API) Initialize() {

}

//Run the API
func (api *API) Run() {
	socket := ":" + api.Port

	log.Println("Starting mango API: " + api.ToJSON())
	log.Fatal(http.ListenAndServe(socket, api.Router))
}

//ToJSON return Json string representation of API struct
func (api *API) ToJSON() string {
	json, _ := json.Marshal(api)
	return string(json)
}
