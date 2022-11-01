package users

import (
	"github.com/4kord/english-flashcards/pkg/services/users"
	"go.uber.org/zap"
)

type Controller struct {
	UsersService users.Service
	Log          *zap.Logger
}

func New(usersService users.Service, log *zap.Logger) *Controller {
	return &Controller{
		UsersService: usersService,
		Log:          log,
	}
}
