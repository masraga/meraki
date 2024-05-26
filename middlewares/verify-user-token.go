package middlewares

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/masraga/meraki/pkg"
	"github.com/masraga/meraki/repositories"
)

func VerifyUserToken(ctx *fiber.Ctx) error {
	autoload := pkg.NewAutoload()
	headers := ctx.GetReqHeaders()
	if len(headers["Authorization"]) == 0 {
		return ctx.Status(http.StatusPreconditionFailed).JSON(fiber.Map{
			"statusCode": http.StatusPreconditionFailed,
			"message":    "unauthorized",
		})
	}
	token := headers["Authorization"][0]
	dToken, err := autoload.JwtHelper().DecodeToken(token)

	/*
		validation token input
	*/
	if err != nil {
		return ctx.Status(http.StatusPreconditionFailed).JSON(fiber.Map{
			"statusCode": http.StatusPreconditionFailed,
			"message":    "unauthorized",
		})
	}

	/*
		check token expiration
	*/
	expiredTime, err := dToken.GetExpirationTime()
	if err != nil {
		return ctx.Status(http.StatusPreconditionFailed).JSON(fiber.Map{
			"statusCode": http.StatusPreconditionFailed,
			"message":    "error parse expired time",
		})
	}

	if time.Now().Unix() > expiredTime.Unix() {
		return ctx.Status(http.StatusPreconditionFailed).JSON(fiber.Map{
			"statusCode": http.StatusPreconditionFailed,
			"message":    "token is expired",
		})
	}

	/*
		make sure user credential is valid
	*/
	userRepo := repositories.NewUser()
	userId, err := dToken.GetIssuer()
	if err != nil {
		return ctx.Status(http.StatusPreconditionFailed).JSON(fiber.Map{
			"statusCode": http.StatusPreconditionFailed,
			"message":    "user not found",
		})
	}
	_, err = userRepo.FindById(userId)
	if err != nil {
		return ctx.Status(http.StatusPreconditionFailed).JSON(fiber.Map{
			"statusCode": http.StatusPreconditionFailed,
			"message":    "user not found",
		})
	}

	return ctx.Next()
}
