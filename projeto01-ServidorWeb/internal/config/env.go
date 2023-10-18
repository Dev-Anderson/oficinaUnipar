package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host         string
	Port         string
	User         string
	Password     string
	NameDatabase string
}

func LoadEnv() (Env, error) {
	godotenv.Load()
	return Env{
		Host:         os.Getenv("api-db-host"),
		Port:         os.Getenv("api-db-port"),
		User:         os.Getenv("api-db-user"),
		Password:     os.Getenv("api-db-password"),
		NameDatabase: os.Getenv("api-db-name"),
	}, nil
}
