package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/masraga/meraki/controllers"
	"github.com/masraga/meraki/middlewares"
)

func Api(router *gin.Engine) {
	/*
		define all controller here
	*/
	userController := controllers.NewUser()
	welcomeController := controllers.NewWelcome()

	/*
		add public api with `api` keyword
	*/
	api := router.Group("/api")
	api.GET("/users/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	api.POST("/users/register", userController.Register)
	api.POST("/users/login", userController.Login)

	/*
		add private api with `apiAuth` keyword
	*/
	apiAuth := api.Use(middlewares.VerifyUserToken)
	apiAuth.GET("/admin/dashboard", welcomeController.Index)
}
