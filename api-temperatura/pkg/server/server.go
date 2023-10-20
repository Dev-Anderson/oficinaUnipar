package server

import (
	"api-temperatura/internal/config"
	"api-temperatura/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type server struct {
	port   string
	server *gin.Engine
}

func NewServer() server {
	env, err := config.LoadEnv()
	if err != nil {
		log.Panicln("Falha ao carregar as variaveis de ambiente ", err.Error())
	}
	return server{
		port:   env.ApiPort,
		server: gin.Default(),
	}
}

func (s *server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Fatal(router.Run(":" + s.port))
}
