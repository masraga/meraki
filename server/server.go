package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/masraga/meraki/pkg"
	"github.com/masraga/meraki/routes"
)

func main() {
	config := pkg.NewConfig("./")

	gin.SetMode(gin.DebugMode)
	router := gin.New()
	routes.Api(router)
	router.Run(fmt.Sprintf(":%s", config.SystemPort))
}
