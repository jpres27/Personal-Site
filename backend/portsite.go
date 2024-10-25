package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
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
	msg := fmt.Sprintf("Blog ID: %d", id)
	w.Write([]byte(msg))
}

func projectsHub(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Projects!"))
}

func decoder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Decoder"))
}

func raycaster(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Raycast rpg"))
}

func dx11engine(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The Isles!"))
}

func libraProject(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Libra contribution!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/blog", blog)
	mux.HandleFunc("/blog/{id}", blogPost)
	mux.HandleFunc("/projects", projectsHub)
	mux.HandleFunc("/projects/decoder", decoder)
	mux.HandleFunc("/projects/raycaster", raycaster)
	mux.HandleFunc("/projects/dx11engine", dx11engine)
	mux.HandleFunc("/projects/libra", libraProject)

	log.Println("Listening on :3000")

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
