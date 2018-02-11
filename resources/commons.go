package resources

import (
	"net/http"

	"github.com/go-chi/chi"
)

// CommonsResources is anonymous type
type CommonsResources struct{}

// Routes is
func (cmm *CommonsResources) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", cmm.List) // GET /todos - read a list of todos

	// r.Route("/{id}", func(r chi.Router) {
	// 	// r.Use(c.TodoCtx) // lets have a todos map, and lets actually load/manipulate
	// 	r.Get("/", c.Get)       // GET /todos/{id} - read a single todo by :id
	// 	r.Put("/", c.Update)    // PUT /todos/{id} - update a single todo by :id
	// 	r.Delete("/", c.Delete) // DELETE /todos/{id} - delete a single todo by :id
	// 	r.Get("/sync", c.Sync)
	// })

	return r
}

// List a
func (cmm *CommonsResources) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todos list of stuff.."))
}
