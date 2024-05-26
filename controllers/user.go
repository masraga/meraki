package controllers

import (
	"github.com/gofiber/fiber/v2"
	userauth "github.com/masraga/meraki/usecase/user-auth"
)

type User struct{}

func (c *User) Register(ctx *fiber.Ctx) error {
	register := userauth.NewRegister(ctx)
	register.Save()
	return nil
}

func (c *User) Login(ctx *fiber.Ctx) error {
	login := userauth.NewLogin(ctx)
	login.Run()
	return nil
}

func NewUser() *User {
	return &User{}
}
