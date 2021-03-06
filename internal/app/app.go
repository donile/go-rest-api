package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct {
	mux *chi.Mux
}

func Create() *App {
	app := App{
		mux: chi.NewRouter(),
	}

	app.mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
	})

	return &app
}

func (a *App) Run() {
	http.ListenAndServe(":3000", a.mux)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
