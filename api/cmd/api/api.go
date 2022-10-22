package main

import (
	"database/sql"

	"github.com/4kord/english-flashcards/api"
	"github.com/4kord/english-flashcards/api/rest"
	"github.com/cloudinary/cloudinary-go/v2"
	"go.uber.org/zap"
)

func main() {
	// logger
	logger, _ := zap.NewDevelopment()
	defer func() {
		if err := logger.Sync(); err != nil {
			panic(err)
		}
	}()

	// parse env
	env, err := api.NewEnv()
	if err != nil {
		logger.Panic("Error reading env variables", zap.Error(err))
	}

	// create db pool
	db, err := sql.Open(env.MainDBDriver, env.MainDBDSN)
	if err != nil {
		logger.Panic("Error opening sql conn", zap.Error(err))
	}
	defer db.Close()

	// ping db
	if err = db.Ping(); err != nil {
		logger.Panic("Error pinging db", zap.Error(err))
	}

	// create cld instance
	cld, err := cloudinary.NewFromParams(env.CldCloud, env.CldKey, env.CldSecret)
	if err != nil {
		logger.Panic("Error creating cld instance", zap.Error(err))
	}

	// server config
	cfg := api.Config{
		Host:       env.APIHost,
		Port:       env.APIPort,
		MainRouter: rest.MainRouter,
		DB:         db,
		Cld:        cld,
		Logger:     logger,
	}

	logger.Info("Starting server",
		zap.String("host", env.APIHost),
		zap.String("port", env.APIPort),
	)

	logger.Fatal("Server has stopped",
		zap.Error(api.NewServer(cfg).Run()),
	)
}
