package main

import (
	"api-temperatura/internal/config"
	"fmt"
)

func main() {
	e, _ := config.LoadEnv()
	fmt.Println(e.ApiPort)
	// database.ConnectDatabase()

	// s := server.NewServer()
	// s.Run()
}
