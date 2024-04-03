package config

import (
	"os"
)

type Env struct {
	Database struct {
		Host     string
		User     string
		Password string
		Name     string
		Port     string
		Type     string
	}
	Log struct {
		Log string
	}
}

func NewEnv() (*Env, error) {
	var cfg Env
	cfg.Database.Host = os.Getenv("DB_HOST")
	cfg.Database.User = os.Getenv("DB_USER")
	cfg.Database.Password = os.Getenv("DB_PASSWORD")
	cfg.Database.Name = os.Getenv("DB_NAME")
	cfg.Database.Port = os.Getenv("DB_PORT")
	cfg.Database.Type = os.Getenv("DB_TYPE")

	cfg.Log.Log = os.Getenv("LOG_FOLDER")

	return &cfg, nil
}
