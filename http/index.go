package http

import (
	"net/http"
)

func (s *Server) showIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>SMA Website </h1>"))
}
