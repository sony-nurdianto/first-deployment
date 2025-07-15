package handler

import (
	"html/template"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles(
		"web/static/index.html",
		"web/static/404.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := template.ExecuteTemplate(w, "index", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
