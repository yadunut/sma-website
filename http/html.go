package http

import (
	"html/template"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/sirupsen/logrus"
)

var (
	assets    http.FileSystem
	templates *template.Template
)

func init() {
	assets = packr.NewBox("./assets")
	templateBox := packr.NewBox("./templates")
	templates = template.New("")
	templateBox.Walk(func(s string, _ packr.File) error {
		template.Must(templates.New(s).Parse(templateBox.String(s)))
		return nil
	})
}

func renderTemplate(w http.ResponseWriter, template string, templateData interface{}) {
	// TODO: Implement request logging

	if err := templates.ExecuteTemplate(w, template, templateData); err != nil {
		logrus.Errorf("error parsing template %s", template)
	}
}
