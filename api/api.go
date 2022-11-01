package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/4kord/english-flashcards/api/internal/controllers"
	"github.com/4kord/english-flashcards/api/internal/middlewares"
	"github.com/4kord/english-flashcards/pkg/cld"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/maker"
	"github.com/4kord/english-flashcards/pkg/services"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Ctx struct {
	Controllers *controllers.Controllers
	Middlewares *middlewares.Middlewares
	Logger      *zap.Logger
}

type Config struct {
	Host   string
	Port   string
	Secret string
	Router func(ctx *Ctx) chi.Router
	DB     *sql.DB
	CLD    *cloudinary.Cloudinary
	Logger *zap.Logger
}

func NewServer(cfg *Config) *http.Server {
	services := services.New(&services.Config{
		Store: maindb.NewStore(cfg.DB),
		Cld:   cld.New(cfg.CLD),
		Maker: maker.NewJWTMaker(cfg.Secret),
	})

	controllers := controllers.New(&controllers.Config{
		Services: services,
		Logger:   cfg.Logger,
	})

	middlewares := middlewares.New(&middlewares.Config{
		Maker:  maker.NewJWTMaker(cfg.Secret),
		Logger: cfg.Logger,
	})

	return &http.Server{
		Addr: cfg.Host + ":" + cfg.Port,
		Handler: cfg.Router(&Ctx{
			Controllers: controllers,
			Middlewares: middlewares,
			Logger:      cfg.Logger,
		}),
		ReadHeaderTimeout: 10 * time.Second,
	}
}
