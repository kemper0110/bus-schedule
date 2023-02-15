package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kemper0110/bus-schedule/internal/models"
)

type DBInterface interface {
	CitiesList(ctx context.Context) ([]models.City, error)
	CarrierList(ctx context.Context) ([]models.Carriers, error)
	SignUp(ctx context.Context, user models.User) (models.Token, error)
	SignIn(ctx context.Context, user models.User) (models.Token, error)
	Schedule(ctx context.Context, reqRoute models.SerchParams) (models.MidRoute, error)
}

type server struct {
	router     *chi.Mux
	httpServer *http.Server
	db         DBInterface
}

type ApiError struct {
	Error string `json:"error"`
}

func newServer(db DBInterface) *server {
	s := &server{
		router: chi.NewRouter(),
		db:     db,
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.Get("/city", s.handlerSearchPage)
	s.router.Post("/sign-up", s.handleSignUp)
	s.router.Post("/sign-in", s.handleSignIn)
	s.router.Get("/schedule", s.handleSchedule)
}

func (s *server) handlerSearchPage(w http.ResponseWriter, r *http.Request) {
	citiesDB, err := s.db.CitiesList(r.Context())
	if err != nil {
		s.httpError(w, r, 400, err)
		return
	}

	s.respond(w, r, 200, citiesDB)
}

func (s *server) handleSchedule(w http.ResponseWriter, r *http.Request) {
	var reqRoute models.SerchParams
	var err error
	reqRoute.FromCity, err = strconv.Atoi(r.URL.Query().Get("from"))
	if err != nil {
		s.httpError(w, r, 400, err)
		return
	}

	reqRoute.ToCity, err = strconv.Atoi(r.URL.Query().Get("to"))
	if err != nil {
		s.httpError(w, r, 400, err)
		return
	}

	reqRoute.FromCity, err = strconv.Atoi(r.URL.Query().Get("date"))
	if err != nil {
		s.httpError(w, r, 400, err)
		return
	}

}

func (s *server) handleSignIn(w http.ResponseWriter, r *http.Request) {
	var reqUser models.User
	reqUser.Login = chi.URLParam(r, "login")
	reqUser.Password = chi.URLParam(r, "password")

	token, err := s.db.SignIn(r.Context(), reqUser)
	if err != nil {
		s.httpError(w, r, 400, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(token.Token))
	fmt.Println(r.Header.Get("Authorization"))
}

func (s *server) handleSignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := s.bodyPars(r, &user); err != nil {
		s.httpError(w, r, 400, err)
		return
	}
	token, err := s.db.SignUp(r.Context(), user)
	if err != nil {
		s.httpError(w, r, 400, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(token.Token))
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) httpError(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, ApiError{Error: err.Error()})
}

func (s *server) bodyPars(r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}
