package google

import (
	"github.com/4kord/english-flashcards/pkg/services/google"
	"go.uber.org/zap"
)

type Controller struct {
	GoogleService google.Service
	Log           *zap.Logger
}

func New(googleService google.Service, log *zap.Logger) *Controller {
	return &Controller{
		GoogleService: googleService,
		Log:           log,
	}
}
