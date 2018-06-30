package http

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// Server is the server struct
type Server struct {
	Port     string
	Logger   logrus.FieldLogger
	Database Database
}

// Listen starts the server.
// This is a blocking function and blocks until server is closed.
func (s *Server) Listen() error {
	s.Logger.Infof("Server starting on port %s", s.Port)
	return http.ListenAndServe(s.Port, s.router())
}

// router creates the router to be served.
func (s *Server) router() http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/register", s.showRegister)
	})

	return r
}
