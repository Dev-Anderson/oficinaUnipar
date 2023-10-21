package controllers

import (
	"api-temperatura/database"
	"api-temperatura/internal/models"
	"api-temperatura/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Criar um novo usuario
// @Summary Cria um novo usuario
// @Description Cria um novo usuario com base nos dados fornecidos
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "Dados do usuário a serem criados"
// @Success 201 {string} 201 "Usuario criado com sucesso"
// @Failure 400 {object} string "Erro de solicitacao: Nao foi possivel vincular o JSON"
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nao foi possivel vincular o JSON" + err.Error(),
		})
		return
	}

	user.Password = services.Sha256Encoder(user.Password)

	err = database.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nao foi possivel criar o usuario" + err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

// Consulta Usuario godoc
// @Summary Obtem a lista de usuarios
// @Description	Retorna a lista de usuarios cadastrados no banco de dados
// @Tags 	User
// @Produce	json
// @Success	200 {array}	models.User
// @Router	/user	[get]
func GetUser(c *gin.Context) {
	var user []models.User
	database.DB.Find(&user)
	c.JSON(http.StatusOK, user)
}
