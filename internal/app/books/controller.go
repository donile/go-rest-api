package books

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	chi.Router
}

func NewController() *Controller {
	controller := Controller{}
	controller.Router = chi.NewRouter()
	controller.Router.Post("/", PostBook)
	return &controller
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
