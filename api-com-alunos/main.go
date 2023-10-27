package main

import (
	"api/database"
	"api/server"
	"fmt"
)

func main() {
	fmt.Println("API Rodando")
	database.Conectabanco()

	s := server.NewServer()
	s.Run()
}
