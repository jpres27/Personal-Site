package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("../frontend/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /projects/{$}", app.projectsHub)
	mux.HandleFunc("GET /projects/raycaster/{$}", app.raycaster)
	mux.HandleFunc("GET /projects/decoder/{$}", app.decoder)

	return mux
}
