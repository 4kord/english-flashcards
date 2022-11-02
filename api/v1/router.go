package v1

import (
	"time"

	"github.com/4kord/english-flashcards/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const userIDParam = "userID"

func Router(c *api.Ctx) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	// public routes
	r.Group(func(r chi.Router) {
		// auth endpoints
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", c.Controllers.Auth.Login)
			r.Post("/register", c.Controllers.Auth.Register)
			r.Post("/logout", c.Controllers.Auth.Logout)
			r.Get("/refresh", c.Controllers.Auth.Refresh)
		})
	})

	// any auth endpoints
	r.Group(func(r chi.Router) {
		r.Use(c.Middlewares.AnyAuth.Handler)

		// users endpoints
		r.Route("/users", func(r chi.Router) {
			r.Route("/{userID}", func(r chi.Router) {
				r.Route("/decks", func(r chi.Router) {
					r.Get("/", c.Controllers.Decks.GetDecks)
					r.Post("/", c.Controllers.Decks.CreateDeck)
				})
			})
		})

		// decks endpoints
		r.Route("/decks", func(r chi.Router) {
			r.Route("/{deckID}", func(r chi.Router) {
				r.Put("/", nil)
				r.Delete("/", nil)

				r.Get("/cards", nil)
				r.Post("/cards", nil)

				r.Post("/insert", nil)
			})

			r.Get("/premade", nil)
		})

		// cards endpoints
		r.Route("/cards", func(r chi.Router) {
			r.Route("/{cardID}", func(r chi.Router) {
				r.Put("/", nil)
				r.Delete("/", nil)
			})
		})

		// google endpoints
		r.Route("/google", func(r chi.Router) {
			r.Route("/audio", nil)
		})
	})

	// admin only endpoints
	r.Group(func(r chi.Router) {
		r.Use(c.Middlewares.AdminAuth.Handler)

		// users endpoints
		r.Route("/users", func(r chi.Router) {
			r.Get("/", nil)

			r.Route("/{userID}", func(r chi.Router) {
				r.Delete("/", nil)

				r.Route("/decks", func(r chi.Router) {
					r.Post("/premade", nil)
				})
			})
		})
	})

	return r
}
