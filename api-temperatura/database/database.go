package database

import (
	"api-temperatura/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Panicln("Falha ao carregar as variaveis de ambiente ", err.Error())
	}

	stringDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.Host, env.Port, env.User, env.Password, env.NameDatabase)

	DB, err = gorm.Open(postgres.Open(stringDB))
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate()

}
