package main

import (
	"api-temperatura/database"
	"api-temperatura/pkg/server"
)

func main() {
	database.ConnectDatabase()

	s := server.NewServer()
	s.Run()
}
