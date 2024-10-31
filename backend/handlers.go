package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.html", templateData{})
}

func (app *application) projectsHub(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "projects-hub.html", templateData{})
}

func (app *application) raycaster(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "raycaster.html", templateData{})
}

func (app *application) decoder(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "decoder.html", templateData{})
}
