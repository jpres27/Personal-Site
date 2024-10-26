package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"../frontend/base.html",
		"../frontend/home.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (app *application) blog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Blog!"))
}

func (app *application) blogPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 || id > 64 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Blog ID: %d", id)
}

func (app *application) projectsHub(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Projects!"))
}

func (app *application) project(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 || id > 8 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Project ID: %d", id)
}
