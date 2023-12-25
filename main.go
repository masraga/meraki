/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/masraga/meraki/cmd"
	pkg "github.com/masraga/meraki/pkg"
)

func main() {
	cmd.Execute()

	config := pkg.NewConfig("./")
	fmt.Print(config.DbName)
}
