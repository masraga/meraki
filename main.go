/*
Copyright © 2023 koderpedia <koderpedia@gmail.com>
*/
package main

import (
	"fmt"
	"time"

	"github.com/masraga/meraki/cmd"
	"github.com/masraga/meraki/pkg"
)

func main() {
	cmd.Execute()

	autoload := pkg.NewAutoload()
	jwt := autoload.JwtHelper()
	token, err := jwt.GenerateToken(time.Now().Add(60*time.Minute), "user")
	if err != nil {
		panic(err)
	}
	decode, err := jwt.DecodeToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Print(decode.GetExpirationTime())
}
