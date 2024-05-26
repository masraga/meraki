package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/masraga/meraki/pkg"
	"github.com/masraga/meraki/routes"
)

func main() {
	env := pkg.NewConfig("./.env")

	app := fiber.New()
	app.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("PONG")
	})

	routes.Api(app)
	app.Listen(fmt.Sprintf(":%s", env.SystemPort))
}
