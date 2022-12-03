package v1

import (
	"time"

	"github.com/4kord/english-flashcards/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
)

const userIDParam = "userID"

func Router(c *api.Ctx) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Test"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	// prometheus
	r.Get("/metrics", promhttp.Handler().ServeHTTP)

	// auth endpoints
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", c.Controllers.Auth.Login)
		r.Post("/register", c.Controllers.Auth.Register)
		r.Post("/logout", c.Controllers.Auth.Logout)
		r.Get("/refresh", c.Controllers.Auth.Refresh)
		r.With(c.Middlewares.AnyAuth.Handler).Get("/user", c.Controllers.Auth.User)
	})

	// users endpoints
	r.Route("/users", func(r chi.Router) {
		r.With(c.Middlewares.AdminAuth.Handler).Get("/", c.Controllers.Users.GetUsers)

		r.Route("/{userID}", func(r chi.Router) {
			r.With(c.Middlewares.AdminAuth.Handler).Delete("/", c.Controllers.Users.DeleteUser)

			r.Route("/decks", func(r chi.Router) {
				r.Get("/", c.Controllers.Decks.GetDecks)
				r.Post("/", c.Controllers.Decks.CreateDeck)

				r.With(c.Middlewares.AdminAuth.Handler).Post("/premade", c.Controllers.Decks.CreatePremadeDeck)
			})
		})
	})

	// decks endpoints
	r.Route("/decks", func(r chi.Router) {
		r.Route("/{deckID}", func(r chi.Router) {
			r.Put("/", c.Controllers.Decks.EditDeck)
			r.Delete("/", c.Controllers.Decks.DeleteDeck)

			r.Get("/cards", c.Controllers.Cards.GetCards)
			r.Post("/cards", c.Controllers.Cards.CreateCard)

			r.Post("/insert", c.Controllers.Cards.InsertCards)
		})

		r.Get("/premade", c.Controllers.Decks.GetPremadeDecks)
	})

	// cards endpoints
	r.Route("/cards", func(r chi.Router) {
		r.Route("/{cardID}", func(r chi.Router) {
			r.Put("/", c.Controllers.Cards.EditCard)
			r.Delete("/", c.Controllers.Cards.DeleteCard)
		})
	})

	// google endpoints
	r.Route("/google", func(r chi.Router) {
		r.Get("/audio/{word}", c.Controllers.Google.FetchAudio)
	})

	return r
}
