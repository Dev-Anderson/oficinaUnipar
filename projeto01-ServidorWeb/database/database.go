package database

import (
	"fmt"
	"log"
	"servidor-web/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Printf("Falha ao carregar o env = %s", err)
	}

	stringDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.Host, env.Port, env.User, env.Password, env.NameDatabase)

	DB, err = gorm.Open(postgres.Open(stringDB))
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate()
}
