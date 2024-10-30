package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.html", templateData{})
}

func (app *application) blog(w http.ResponseWriter, r *http.Request) {
	posts, err := app.posts.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Posts: posts,
	}

	app.render(w, r, http.StatusOK, "blog.html", data)
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

	data := templateData{
		Post: post,
	}

	app.render(w, r, http.StatusOK, "view-post.html", data)
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
	app.render(w, r, http.StatusOK, "projects-hub.html", templateData{})
}

func (app *application) raycasterIntro(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "raycaster-intro.html", templateData{})
}

func (app *application) decoder(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "decoder.html", templateData{})
}
