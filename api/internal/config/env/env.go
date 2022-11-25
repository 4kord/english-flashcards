package env

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Env struct {
	APIHost       string `env:"API_HOST"`
	APIPort       string `env:"API_PORT"`
	MainDBDriver  string `env:"MAINDB_DRIVER"`
	MainDBDSN     string `env:"MAINDB_DSN"`
	CloudinaryURL string `env:"CLOUDINARY_URL"`
	Secret        string `env:"SECRET"`
}

func Parse() (*Env, error) {
	var v Env

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	err = env.Parse(&v)
	if err != nil {
		return nil, err
	}

	return &v, err
}
