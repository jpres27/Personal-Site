package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("../frontend/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /blog/{$}", app.blog)
	mux.HandleFunc("GET /blog/{id}", app.blogPost)
	mux.HandleFunc("GET /blog/create", app.blogCreateForm)
	mux.HandleFunc("POST /blog/create", app.blogCreatePost)
	mux.HandleFunc("GET /projects/{$}", app.projectsHub)
	mux.HandleFunc("GET /projects/{id}", app.project)

	return mux
}
