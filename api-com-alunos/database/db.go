package database

import (
	"api/config"
	"api/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Conectabanco() error {
	e := config.LoadEnv()

	fmt.Print(e.Host)

	stringDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		e.Host, e.Port, e.User, e.Password, e.DBName)

	DB, err = gorm.Open(postgres.Open(stringDB))
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Teste{})

	return nil
}
