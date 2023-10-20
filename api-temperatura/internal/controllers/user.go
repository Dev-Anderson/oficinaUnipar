package controllers

import (
	"api-temperatura/database"
	"api-temperatura/internal/models"
	"api-temperatura/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func GetUser(c *gin.Context) {
	var user []models.User
	database.DB.Find(&user)
	c.JSON(http.StatusOK, user)
}
