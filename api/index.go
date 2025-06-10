package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("public", "template.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}
