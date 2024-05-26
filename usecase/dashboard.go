package usecase

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Dashboard struct {
	Ctx *fiber.Ctx
}

func (d *Dashboard) Index() {
	d.Ctx.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": http.StatusOK,
		"message":    "welcome to dashboard with authentication",
	})
}

func NewDashboard(ctx *fiber.Ctx) *Dashboard {
	return &Dashboard{
		Ctx: ctx,
	}
}
