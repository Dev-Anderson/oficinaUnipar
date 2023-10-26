package config

import (
	"os"
)

type Env struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadEnv() Env {
	err := LoadEnvFromFile(".env")
	if err != nil {
		panic(err)
	}
	return Env{
		Host:     os.Getenv("API-HOST"),
		Port:     os.Getenv("API-PORT"),
		User:     os.Getenv("API-USER"),
		Password: os.Getenv("API-PASSWORD"),
		DBName:   os.Getenv("API-DBNAME"),
	}

}
