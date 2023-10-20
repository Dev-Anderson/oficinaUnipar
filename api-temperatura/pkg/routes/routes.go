package routes

import (
	"api-temperatura/internal/controllers"

	"api-temperatura/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"

	main := router.Group("api/v1")
	{
		home := main.Group("home")
		{
			home.GET("/", controllers.Home)
		}
		user := main.Group("user")
		{
			user.GET("/", controllers.GetUser)
			user.POST("/", controllers.CreateUser)
		}
		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
		temperatura := main.Group("temperatura")
		{
			temperatura.POST("/", controllers.CreateTemperatura)
			temperatura.GET("/", controllers.GetTemperatura)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
