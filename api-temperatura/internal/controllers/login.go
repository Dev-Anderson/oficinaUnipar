package controllers

import (
	"api-temperatura/database"
	"api-temperatura/internal/models"
	"api-temperatura/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Efetua Login de usuario
// @Description Efetua o login de um usuario com base nos dados fornecidos
// @Tags Login
// @Accept json
// @Produce json
// @Param user body models.Login true "Dados de login do usuario"
// @Success 200 {string} string "Token gerado com sucesso"
// @Failure 400 {object} string "Erro de solicitacao: Nao foi possivel enviar o body"
// @Failure 401 {object} string "Erro interno do servidor: Erro ao gerar o token"
// @Failure 500 {object} string "Erro interno do servidor: Erro ao gerar o token"
// @Router /login [post]
func Login(c *gin.Context) {
	var login models.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erorr": "Nao foi possivel enviar o body" + err.Error(),
		})
		return
	}

	var user models.User
	dbError := database.DB.Where("email = ?", login.Email).First(&user).Error
	if dbError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Falha ao carregar o usuario" + dbError.Error(),
		})
		return
	}

	if user.Password != services.Sha256Encoder(login.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"erorr": "Senha incorreto",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao gerar o token" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": token,
	})
}
