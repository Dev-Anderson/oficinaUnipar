package database

import (
	"api/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Conectabanco() error {
	host := "localhost"
	port := "5432"
	user := "postgres"
	pass := "postgres"
	dbname := "teste"

	stringDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)

	DB, err = gorm.Open(postgres.Open(stringDB))
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Teste{}, &models.Temperatura{})

	return nil
}
