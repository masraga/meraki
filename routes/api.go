package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masraga/meraki/controllers"
	"github.com/masraga/meraki/middlewares"
)

func Api(router *fiber.App) {
	/*
		define all controller here
	*/
	userController := controllers.NewUser()
	welcomeController := controllers.NewWelcome()

	/*
		add public api with `api` keyword
	*/
	api := router.Group("/api")
	api.Get("/users/hello-world", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "your're very cool",
			"status":  200,
		})
	})
	api.Post("/users/register", userController.Register)
	api.Post("/users/login", userController.Login)

	// /*
	// 	add private api with `apiAuth` keyword
	// */
	apiAuth := api.Use(middlewares.VerifyUserToken)
	apiAuth.Get("/admin/dashboard", welcomeController.Index)
}
