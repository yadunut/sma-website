package http

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gobuffalo/packr"
	"github.com/sirupsen/logrus"
	"github.com/yadunut/sma-website/http/middleware"
)

// Server is the server struct
type Server struct {
	Port      string
	Logger    logrus.FieldLogger
	Database  Database
	templates *template.Template
	assets    http.FileSystem
}

func NewServer(Port string, Logger logrus.FieldLogger, Database Database) (*Server, error) {
	s := Server{
		Port:     Port,
		Logger:   Logger,
		Database: Database,
	}
	assets := packr.NewBox("./assets")
	templateBox := packr.NewBox("./templates")

	templates := template.New("")
	templateBox.Walk(func(s string, _ packr.File) error {
		template.Must(templates.New(s).Parse(templateBox.String(s)))
		return nil
	})
	s.assets = assets
	s.templates = templates

	return &s, nil
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
		FileServer(r, "/public", s.assets)
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
