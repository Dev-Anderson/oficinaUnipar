package controllers

import (
	"api-temperatura/database"
	"api-temperatura/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Cadastro de nova temperatura
// @Summary Cadastra uma nova temperatura
// @Description Cadastra uma nova temperatura
// @Tags Temperatura
// @Accept json
// @Produce json
// @Param temperatura body models.Temperatura true "Cadatro de uma nova temperatura"
// @Success 200 {object} models.Temperatura
// @Failure 400 {string} string
// @Router /temperatura [post]
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

// Consulta temperatura godoc
// @Summary Consulta a temperatura
// @Description Retorna a lista de temperatura cadastradas
// @Tags Temperatura
// @Produce json
// @Success 200 {array} models.Temperatura
// @Router /temperatura [get]
func GetTemperatura(c *gin.Context) {
	var temperaturas []models.Temperatura
	database.DB.Find(&temperaturas)

	var response []map[string]interface{}
	for _, temperatura := range temperaturas {
		dataFormatada := temperatura.Data.Format("2006-01-02")

		temperaturaMap := map[string]interface{}{
			"temperatura": temperatura.Temperatura,
			"data":        dataFormatada,
		}

		response = append(response, temperaturaMap)
	}

	c.JSON(http.StatusOK, response)
}
