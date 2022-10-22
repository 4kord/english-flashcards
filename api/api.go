package api

import (
	"database/sql"
	"net/http"

	"github.com/4kord/english-flashcards/pkg/cld"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/services"
	"github.com/caarlos0/env"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type envoirnment struct {
	APIHost      string `env:"API_HOST"`
	APIPort      string `env:"API_PORT"`
	MainDBDriver string `env:"MAINDB_DRIVER"`
	MainDBDSN    string `env:"MAINDB_DSN"`
	CldCloud     string `env:"CLD_CLOUD"`
	CldKey       string `env:"CLD_KEY"`
	CldSecret    string `env:"CLD_SECRET"`
}

func NewEnv() (*envoirnment, error) {
	var envVars envoirnment

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

type MainRouter func(services *services.Services, logger *zap.Logger) chi.Router

type Server struct {
	*http.Server
}

func NewServer(cfg Config) *Server {
	services := services.New(services.Config{
		Store: maindb.NewStore(cfg.DB),
		Cld:   cld.New(cfg.Cld),
	})

	return &Server{
		Server: &http.Server{
			Addr:    cfg.Host + ":" + cfg.Port,
			Handler: cfg.MainRouter(services, cfg.Logger),
		},
	}
}

func (s *Server) Run() error {
	return s.ListenAndServe()
}
