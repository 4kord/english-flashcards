package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/4kord/english-flashcards/pkg/cld"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/services"
	"github.com/caarlos0/env"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Envoirnment struct {
	APIHost      string `env:"API_HOST"`
	APIPort      string `env:"API_PORT"`
	MainDBDriver string `env:"MAINDB_DRIVER"`
	MainDBDSN    string `env:"MAINDB_DSN"`
	CldCloud     string `env:"CLD_CLOUD"`
	CldKey       string `env:"CLD_KEY"`
	CldSecret    string `env:"CLD_SECRET"`
}

func NewEnv() (*Envoirnment, error) {
	var envVars Envoirnment

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	err = env.Parse(&envVars)
	if err != nil {
		return nil, err
	}

	return &envVars, err
}

type Config struct {
	Host string
	Port string
	//
	MainRouter MainRouter
	DB         *sql.DB
	Cld        *cloudinary.Cloudinary
	Logger     *zap.Logger
}

type MainRouter func(s *services.Services, logger *zap.Logger) chi.Router

type Server struct {
	*http.Server
}

func NewServer(cfg Config) *Server {
	s := services.New(&services.Config{
		Store: maindb.NewStore(cfg.DB),
		Cld:   cld.New(cfg.Cld),
	})

	return &Server{
		Server: &http.Server{
			Addr:              cfg.Host + ":" + cfg.Port,
			Handler:           cfg.MainRouter(s, cfg.Logger),
			ReadHeaderTimeout: 2 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	return s.ListenAndServe()
}
