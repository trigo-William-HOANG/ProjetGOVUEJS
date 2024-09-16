package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func Handler(r *chi.Mux) {
	// Configure CORS options
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Replace with your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Other middlewares
	r.Use(chimiddle.StripSlashes)

	r.Route("/application-data", func(router chi.Router) {
		// router.Use(middleware.AuthMiddleware) // Uncomment if you add authentication middleware
		router.Get("/", GetAppDataHandler)
	})
}
