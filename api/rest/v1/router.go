package v1

import (
	"github.com/4kord/english-flashcards/pkg/services"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Router(r chi.Router, services *services.Services, logger *zap.Logger) chi.Router {
	authcontroller := &authController{
		authService: services.Auth,
		log:         logger,
	}

	cardscontroller := &cardsController{
		cardsService: services.Cards,
		log:          logger,
	}

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authcontroller.Register) // TODO: EMAIL VERIFICATION
		r.Post("/login", authcontroller.Login)       // TODO: IP SESSION
	})

	r.Route("/cards", func(r chi.Router) {
		r.Get("/", cardscontroller.GetCards)
	})

	r.Route("/decks", func(r chi.Router) {
		r.Post("/{deckID}", cardscontroller.CreateCard)
	})

	return r
}
