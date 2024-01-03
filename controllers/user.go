package controllers

import (
	"github.com/gin-gonic/gin"
	userauth "github.com/masraga/meraki/usecase/user-auth"
)

type User struct{}

func (c *User) Register(ctx *gin.Context) {
	register := userauth.NewRegister(ctx)
	register.Save()
}

func NewUser() *User {
	return &User{}
}
