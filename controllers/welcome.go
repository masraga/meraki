package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masraga/meraki/usecase"
)

type Welcome struct{}

/*
default method for every controller as a reference
for every method in controller
*/
func (c *Welcome) Index(ctx *fiber.Ctx) error {
	dashboard := usecase.NewDashboard(ctx)
	dashboard.Index()

	return nil
}

func NewWelcome() *Welcome {
	return &Welcome{}
}
