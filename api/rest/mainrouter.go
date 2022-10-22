package rest

import (
	v1 "github.com/4kord/english-flashcards/api/rest/v1"
	"github.com/4kord/english-flashcards/pkg/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

func MainRouter(services *services.Services, logger *zap.Logger) chi.Router {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.AllowAll().Handler)

	r.Route("/v1", func(r chi.Router) {
		v1.Router(r, services, logger)
	})

	return r
}
