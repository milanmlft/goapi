package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/milanmlft/goapi/internal/middleware"
)

// Exported function to set up the Chi router
// This sets up the router's endpoints and any middleware
func Handler(r *chi.Mux) {
	// Global middleware to strip trailing slashes from endpoints
	r.Use(chimiddle.StripSlashes)

	// Set up store endpoint with a GET method to retrieve ice cream flavours
	r.Route("/account", func(router chi.Router) {
		// Middleware to check for authorisation
		router.Use(middleware.Authorisation)
		router.Get("/flavours", GetFlavours)
	})
}
