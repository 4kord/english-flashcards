package middlewares

import (
	"github.com/4kord/english-flashcards/pkg/maker"
	"go.uber.org/zap"
)

type Config struct {
	Maker  maker.Maker
	Logger *zap.Logger
}

type Middlewares struct {
	AdminAuth *Auth
	AnyAuth   *Auth
}

func New(cfg *Config) *Middlewares {
	return &Middlewares{
		AnyAuth:   NewAuth(cfg.Maker, cfg.Logger, false),
		AdminAuth: NewAuth(cfg.Maker, cfg.Logger, true),
	}
}
