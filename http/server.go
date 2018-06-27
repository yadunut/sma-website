package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"mod/github.com/sirupsen/logrus@v1.0.5"
)

// Server ...
type Server struct {
	Port     string
	Logger   logrus.FieldLogger
	Database Database
}

// Listen ...
func (s *Server) Listen() error {
	return http.ListenAndServe(s.Port, s.router())
}

// router creates the router that
func (s *Server) router() http.Handler {
	r := chi.NewRouter()
	return r
}
