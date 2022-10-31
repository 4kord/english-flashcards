package auth

import (
	"github.com/4kord/english-flashcards/pkg/services/auth"
	"go.uber.org/zap"
)

type Controller struct {
	AuthService auth.Service
	Log         *zap.Logger
}

func New(authService auth.Service, log *zap.Logger) *Controller {
	return &Controller{
		AuthService: authService,
		Log:         log,
	}
}
