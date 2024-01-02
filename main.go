/*
Copyright © 2023 koderpedia <koderpedia@gmail.com>
*/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/masraga/meraki/cmd"
	"github.com/masraga/meraki/pkg"
	"github.com/masraga/meraki/routes"
)

func main() {
	cmd.Execute()

	config := pkg.NewConfig("./")

	router := gin.Default()
	routes.Api(router)
	router.Run(fmt.Sprintf(":%s", config.SystemPort))
}
