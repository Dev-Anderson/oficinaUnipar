package main

import (
	"api-temperatura/internal/config"
	"fmt"
)

func main() {
	fmt.Println("Carregando as variaveis de ambiente")
	env, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	fmt.Println("host: ", env.Host)
	fmt.Println("port: ", env.Port)
}
