package app

import (
	"encoding/json"
	"net/http"

	"github.com/IB133/bus-tickets/internal/db"
	"github.com/go-chi/chi/v5"
)

type server struct {
	router     *chi.Mux
	httpServer *http.Server
	db         *db.DB
}

func newServer(dbConnURI string) *server {
	return &server{
		router: chi.NewRouter(),
	}
}

func (s *server) configureRouter() {
	s.router.Get("/", s.handlerSearchPage)
}

func (s *server) handlerSearchPage(w http.ResponseWriter, r *http.Request) {

}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	http.Error(w, err.Error(), code)
}
