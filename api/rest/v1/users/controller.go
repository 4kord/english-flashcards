package users

import (
	"github.com/4kord/english-flashcards/pkg/services/users"
	"go.uber.org/zap"
)

type Controller struct {
	UserService users.Service
	Log         *zap.Logger
}
