package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/yadunut/sma-website/http/middleware"
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
		r.Get("/", s.showIndex)
		r.Get("/register", s.showRegister)
		FileServer(r, "/public", assets)
	})

	return r
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.With(middleware.Neuter).Get(path, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}
