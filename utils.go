package main

import (
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	templates := template.Must(template.ParseGlob("templates/*"))
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
