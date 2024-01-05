package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/masraga/meraki/controllers"
	"github.com/masraga/meraki/middlewares"
)

func Api(router *gin.Engine) {

	userController := controllers.NewUser()
	welcomeController := controllers.NewWelcome()

	api := router.Group("/api")

	api.GET("/users/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	api.POST("/users/register", userController.Register)
	api.POST("/users/login", userController.Login)

	apiAuth := api.Use(middlewares.VerifyUserToken)
	apiAuth.GET("/admin/dashboard", welcomeController.Index)
}
