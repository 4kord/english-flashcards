package setup

import (
	"database/sql"

	"github.com/4kord/english-flashcards/api/internal/config/env"
	"github.com/cloudinary/cloudinary-go/v2"
	"go.uber.org/zap"
)

type App struct {
	DB     *sql.DB
	CLD    *cloudinary.Cloudinary
	Secret string
	Logger *zap.Logger
	Config *ServerConfig
}

func New() *App {
	var app = new(App)

	vars, _ := env.Parse()

	app.DB = setupDB(vars.MainDBDriver, vars.MainDBDSN)

	app.CLD = setupCLD(vars.CloudinaryURL)

	app.Secret = vars.Secret

	app.Logger = setupLogger()

	app.Config = &ServerConfig{
		Host: vars.APIHost,
		Port: vars.APIPort,
	}

	return app
}
