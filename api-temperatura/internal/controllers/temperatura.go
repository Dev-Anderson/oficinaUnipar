package controllers

import (
	"api-temperatura/database"
	"api-temperatura/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTemperatura(c *gin.Context) {
	var temperatura models.Temperatura
	if err := c.ShouldBindJSON(&temperatura); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	temperatura.Data = time.Now()

	database.DB.Create(&temperatura)
	c.JSON(http.StatusOK, temperatura)
}
