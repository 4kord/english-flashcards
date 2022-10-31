package main

import (
	"github.com/4kord/english-flashcards/api"
	"github.com/4kord/english-flashcards/api/internal/config/setup"
	v1 "github.com/4kord/english-flashcards/api/v1"
	"go.uber.org/zap"
)

func main() {
	app := setup.New()

	defer app.DB.Close()

	defer func() {
		if err := app.Logger.Sync(); err != nil {
			panic(err)
		}
	}()

	api := api.NewServer(&api.Config{
		Host:   app.Config.Host,
		Port:   app.Config.Port,
		Secret: app.Secret,
		Router: v1.Router,
		DB:     app.DB,
		CLD:    app.CLD,
		Logger: app.Logger,
	})

	app.Logger.Info("Server is running", zap.String("host", app.Config.Host))

	_ = api.ListenAndServe()
}
