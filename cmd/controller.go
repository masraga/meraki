/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/masraga/meraki/pkg"
	"github.com/spf13/cobra"
)

// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "create controller",
	Long:  `create controller`,
	Run: func(cmd *cobra.Command, args []string) {
		autoload := pkg.NewAutoload()
		ctlName, _ := cmd.Flags().GetString("name")
		ctlFileName := fmt.Sprintf("./controllers/%s", autoload.FilenameFormatHelper(ctlName, "go"))

		var ctlScript []string

		ctlScript = append(ctlScript, "package controllers\n\n")

		ctlScript = append(ctlScript, "import (\n")
		ctlScript = append(ctlScript, "\t\"fmt\"\n\n")
		ctlScript = append(ctlScript, "\t\"github.com/gin-gonic/gin\"\n")
		ctlScript = append(ctlScript, ")\n\n")

		ctlScript = append(ctlScript, fmt.Sprintf("type %s struct{}\n\n", ctlName))

		ctlScript = append(ctlScript, "/*\n")
		ctlScript = append(ctlScript, "default method for every controller as a reference\n")
		ctlScript = append(ctlScript, "for every method in controller\n")
		ctlScript = append(ctlScript, "*/\n")
		ctlScript = append(ctlScript, fmt.Sprintf("func (c *%s) Index(ctx *gin.Context) {\n", ctlName))
		ctlScript = append(ctlScript, fmt.Sprintf("\tfmt.Print(\"index method is called in controller %s\")\n", ctlName))
		ctlScript = append(ctlScript, "}\n\n")

		ctlScript = append(ctlScript, fmt.Sprintf("func New%s() *%s {\n", ctlName, ctlName))
		ctlScript = append(ctlScript, fmt.Sprintf("\treturn &%s{}\n", ctlName))
		ctlScript = append(ctlScript, "}")

		if _, err := os.Stat(ctlFileName); errors.Is(err, os.ErrNotExist) {
			os.WriteFile(ctlFileName, []byte(strings.Join(ctlScript, "")), 0664)
			fmt.Println(`controller created in:`, ctlFileName)
		} else {
			panic(fmt.Errorf("[error-log] controller %s is exists", ctlFileName))
		}
	},
}

func init() {
	rootCmd.AddCommand(controllerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	controllerCmd.PersistentFlags().String("name", "", "controller name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// controllerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
