package v1

import (
	"github.com/4kord/english-flashcards/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(c *api.Ctx) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", c.Controllers.Auth.Register) // TODO: EMAIL VERIFICATION
			r.Post("/login", c.Controllers.Auth.Login)       // TODO: IP SESSION
			r.Post("/logout", c.Controllers.Auth.Logout)
		})

		r.Route("/users", func(r chi.Router) {
			r.Use(c.Middlewares.Auth.Handler)

			r.Get("/", c.Controllers.Users.GetUsers)
			r.Delete("/{userID:[0-9]+}", c.Controllers.Users.DeleteUser) // TODO: Remove FK Constraint

			r.Route("/{userID:[0-9]+}/decks", func(r chi.Router) {
				r.Get("/", c.Controllers.Decks.GetDecks)
				r.Post("/", c.Controllers.Decks.CreateDeck)

				r.Post("/premade", c.Controllers.Decks.CreatePremadeDeck)
			})
		})

		r.Route("/decks", func(r chi.Router) {
			r.Route("/{deckID:[0-9]+}", func(r chi.Router) {
				r.Get("/", c.Controllers.Cards.GetCards)
				r.Post("/", c.Controllers.Cards.CreateCard)
				r.Put("/", c.Controllers.Decks.EditDeck)
				r.Delete("/", c.Controllers.Decks.DeleteDeck)

				r.Post("/insert", c.Controllers.Cards.InsertCards)
			})

			r.Get("/premade", c.Controllers.Decks.GetPremadeDecks)
		})

		r.Route("/cards", func(r chi.Router) {
			r.Route("/{cardID:[0-9]+}", func(r chi.Router) {
				r.Put("/", c.Controllers.Cards.EditCard)
				r.Delete("/", c.Controllers.Cards.DeleteCard)
			})
		})

		r.Route("/google", func(r chi.Router) {
			r.Get("/audio/{word:[a-z]+}", c.Controllers.Google.FetchAudio)
		})
	})

	return r
}
