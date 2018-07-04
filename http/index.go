package http

import (
	"net/http"
)

func (s *Server) showIndex(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.tmpl", nil)
}
