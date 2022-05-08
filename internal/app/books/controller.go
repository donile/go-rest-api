package books

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Controller struct {
	chi.Router
	repository *Repository
}

func NewController(r *Repository) *Controller {
	controller := Controller{
		Router:     chi.NewRouter(),
		repository: r,
	}
	controller.Router.Post("/", controller.PostBook)
	return &controller
}

func (c *Controller) PostBook(w http.ResponseWriter, r *http.Request) {
	book := &Book{}

	if err := render.Bind(r, book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	c.repository.Add(book)

	w.WriteHeader(http.StatusCreated)
}
