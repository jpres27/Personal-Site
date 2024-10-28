package main

import (
	"errors"
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
		app.serverError(w, r, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) blog(w http.ResponseWriter, r *http.Request) {
	posts, err := app.posts.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, post := range posts {
		fmt.Fprintf(w, "%+v", post)
	}
}

func (app *application) blogPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 || id > 64 {
		http.NotFound(w, r)
		return
	}

	post, err := app.posts.Get(id)
	if err != nil {
		if errors.Is(err, ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", post)
}

func (app *application) blogCreateForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a page for writing and uploading a blog post"))
}

func (app *application) blogCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "Blogging from the app"
	content := "This blog post was created from the app"

	id, err := app.posts.Insert(title, content)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/blog/%d", id), http.StatusSeeOther)
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
