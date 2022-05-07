package app

import (
	"app/internal/app/books"
	"app/internal/app/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct {
	mux      *chi.Mux
	services *services.Container
}

func Create() *App {
	app := App{
		mux: chi.NewRouter(),
	}

	builder := services.NewBuilder()
	builder.AddSingleton("bookController", func() (interface{}, error) { return books.NewController(), nil })

	app.services = builder.Build()

	bookController := app.services.GetRequiredService("bookController").(*books.Controller)
	app.mux.Mount("/books", bookController)

	return &app
}

func (a *App) Run() {
	http.ListenAndServe(":3000", a.mux)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
