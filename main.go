/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/masraga/meraki/cmd"
	"github.com/masraga/meraki/routes"
)

func main() {
	cmd.Execute()

	router := gin.Default()
	routes.Api(router)
	router.Run(":8003")
}
