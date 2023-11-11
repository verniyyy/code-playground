package src

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Router ...
type Router struct {
	*chi.Mux
}

// NewRouter ...
func NewRouter() Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := Handler{}
	r.Route("/run", func(r chi.Router) {
		r.Post("/go", h.RunGo)
		r.Post("/ruby", h.RunRuby)
	})

	return Router{r}
}
