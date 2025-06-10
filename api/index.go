package handler

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Support local + Vercel by detecting the env var
	publicDir := os.Getenv("VERCEL_PUBLIC_DIR")
	if publicDir == "" {
		publicDir = "public"
	}

	tmplPath := filepath.Join(publicDir, "template.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}
