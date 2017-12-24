package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// API propeties
type API struct {
	Name    string
	Version *Version
	Port    string
	Router  chi.Router
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
	log.Fatal(http.ListenAndServe(socket, api.Router))
}
