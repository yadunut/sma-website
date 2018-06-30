package http

import (
	"net/http"
)

func (s *Server) showRegister(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "register.tmpl", nil)
}
