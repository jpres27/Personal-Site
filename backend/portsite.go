package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"../frontend/base.html",
		"../frontend/home.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func blog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Blog!"))
}

func blogPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 || id > 64 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Blog ID: %d", id)
}

func projectsHub(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Projects!"))
}

func project(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 || id > 8 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Project ID: %d", id)
}

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("../frontend/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /blog", blog)
	mux.HandleFunc("GET /blog/{id}", blogPost)
	mux.HandleFunc("GET /projects", projectsHub)
	mux.HandleFunc("GET /projects/{id}", project)

	log.Println("Listening on :3000")

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
