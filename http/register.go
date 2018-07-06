package http

import (
	"net/http"
)

func (s *Server) showRegister(w http.ResponseWriter, r *http.Request) {
	s.renderTemplate(w, "register.tmpl", nil)
}
