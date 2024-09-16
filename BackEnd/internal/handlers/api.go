package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)
	r.Route("/application-data", func(router chi.Router) {
		// router.Use(middleware.AuthMiddleware) au cas ou je vais rajouter une sécurité
		router.Get("/", GetAppDataHandler)
	})
}
