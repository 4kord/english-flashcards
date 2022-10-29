package v1

import (
	"github.com/4kord/english-flashcards/api/rest/v1/auth"
	"github.com/4kord/english-flashcards/api/rest/v1/cards"
	"github.com/4kord/english-flashcards/api/rest/v1/decks"
	"github.com/4kord/english-flashcards/api/rest/v1/google"
	"github.com/4kord/english-flashcards/api/rest/v1/users"
	"github.com/4kord/english-flashcards/pkg/services"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Router(r chi.Router, s *services.Services, log *zap.Logger) chi.Router {
	authcontroller := &auth.Controller{
		AuthService: s.Auth,
		Log:         log,
	}

	userController := &users.Controller{
		UserService: s.Users,
		Log:         log,
	}

	cardscontroller := &cards.Controller{
		CardService: s.Cards,
		Log:         log,
	}

	deckController := &decks.Controller{
		DeckService: s.Decks,
		Log:         log,
	}

	googleController := &google.Controller{
		GoogleService: s.Google,
		Log:           log,
	}

	r.Get("/user", authcontroller.User)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authcontroller.Register) // TODO: EMAIL VERIFICATION
		r.Post("/login", authcontroller.Login)       // TODO: IP SESSION
		r.Post("/logout", authcontroller.Logout)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userController.GetUsers)
		r.Delete("/{userID:[0-9]+}", userController.DeleteUser) // TODO: Remove FK Constraint

		r.Route("/{userID:[0-9]+}/decks", func(r chi.Router) {
			r.Get("/", deckController.GetDecks)
			r.Post("/", deckController.CreateDeck)

			r.Post("/premade", deckController.CreatePremadeDeck)
		})
	})

	r.Route("/decks", func(r chi.Router) {
		r.Route("/{deckID:[0-9]+}", func(r chi.Router) {
			r.Get("/", cardscontroller.GetCards)
			r.Post("/", cardscontroller.CreateCard)
			r.Put("/", deckController.EditDeck)
			r.Delete("/", deckController.DeleteDeck)

			r.Post("/insert", cardscontroller.InsertCards)
		})

		r.Get("/premade", deckController.GetPremadeDecks)
	})

	r.Route("/cards", func(r chi.Router) {
		r.Route("/{cardID:[0-9]+}", func(r chi.Router) {
			r.Put("/", cardscontroller.EditCard)
			r.Delete("/", cardscontroller.DeleteCard)
		})
	})

	r.Route("/google", func(r chi.Router) {
		r.Get("/audio/{word:[a-z]+}", googleController.FetchAudio)
	})

	return r
}
