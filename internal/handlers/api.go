package handlers

import (
	"github.com/MoNezhadali/go_api/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler (r *chi.Mux) {
	// Middleware is a function a function that is called before the primary function handles the endpoint.
	// you can add middleware to a rout by .Use function
	// Global middleware (It is applied all the time)
	// All StripSlashes middleware will remove the trailing slash from the URL path, so /account/ will be treated as /account
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		// Middleware for /account route
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}