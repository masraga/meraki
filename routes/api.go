package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/masraga/meraki/controllers"
)

func Api(router *gin.Engine) {

	userController := controllers.NewUser()

	api := router.Group("/api")

	api.GET("/users/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	api.POST("/users/register", userController.Register)
	api.POST("/users/login", userController.Login)
}
