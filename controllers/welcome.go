package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/masraga/meraki/usecase"
)

type Welcome struct{}

/*
default method for every controller as a reference
for every method in controller
*/
func (c *Welcome) Index(ctx *gin.Context) {
	dashboard := usecase.NewDashboard(ctx)
	dashboard.Index()
}

func NewWelcome() *Welcome {
	return &Welcome{}
}
