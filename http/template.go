package http

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

var ()

func init() {
}

func (s Server) renderTemplate(w http.ResponseWriter, template string, templateData interface{}) {

	if err := s.templates.ExecuteTemplate(w, template, templateData); err != nil {
		logrus.Errorf("error parsing template %s", template)
	}
}
