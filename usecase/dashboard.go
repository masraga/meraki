package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Dashboard struct {
	Ctx *gin.Context
}

func (d *Dashboard) Index() {
	d.Ctx.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "welcome to dashboard with authentication",
	})
}

func NewDashboard(ctx *gin.Context) *Dashboard {
	return &Dashboard{
		Ctx: ctx,
	}
}
