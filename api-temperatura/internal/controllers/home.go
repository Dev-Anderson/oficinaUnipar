package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home godoc
// @Summary Retorna uma mensagem de boas-vindas da API
// @Description Retorna uma mensagem de boas-vindas da API Temperatura
// @Tags 	Bem-vindo
// @Produce  json
// @Success 200 {string} string "message: Api Temperatura rodando com sucesso!"
// @Router /home [get]
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Api Temperatura rodando com sucesso!",
	})
}
