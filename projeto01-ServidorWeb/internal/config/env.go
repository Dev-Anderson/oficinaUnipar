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
		Host:         os.Getenv("api_db_host"),
		Port:         os.Getenv("api_db_port"),
		User:         os.Getenv("api_db_user"),
		Password:     os.Getenv("api_db_password"),
		NameDatabase: os.Getenv("api_db_name"),
	}, nil
}
