package app

import (
	"app/internal/app/books"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

type App struct {
	mux *chi.Mux
}

func Create() *App {
	app := App{
		mux: chi.NewRouter(),
	}

	psqlInfo := "host=localhost port=5432 user=postgres password=password dbname=book sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	bookRepository := books.NewRepository(db)
	bookController := books.NewController(bookRepository)
	app.mux.Mount("/books", bookController)

	return &app
}

func (a *App) Run() {
	http.ListenAndServe(":3000", a.mux)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
