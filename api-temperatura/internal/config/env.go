package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	Host         string
	Port         string
	User         string
	Password     string
	NameDatabase string
	ApiPort      string
	SecretKey    string
	Issure       string
	TimeExpires  string
}

func LoadEnv() (ConfigEnv, error) {
	godotenv.Load()
	return ConfigEnv{
		Host:         os.Getenv("db_host"),
		Port:         os.Getenv("db_port"),
		User:         os.Getenv("db_user"),
		Password:     os.Getenv("db_password"),
		NameDatabase: os.Getenv("db_name"),
		ApiPort:      os.Getenv("api_port"),
		SecretKey:    os.Getenv("secretKey"),
		Issure:       os.Getenv("issure"),
		TimeExpires:  os.Getenv("timeExpires"),
	}, nil
}
