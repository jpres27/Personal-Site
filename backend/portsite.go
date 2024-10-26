package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("../frontend/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /blog", app.blog)
	mux.HandleFunc("GET /blog/{id}", app.blogPost)
	mux.HandleFunc("GET /projects", app.projectsHub)
	mux.HandleFunc("GET /projects/{id}", app.project)

	logger.Info("starting server", slog.String("addr", ":3000"))
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
