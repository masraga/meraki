package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/masraga/meraki/pkg"
	"github.com/masraga/meraki/repositories"
)

func VerifyUserToken(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	jwt := pkg.NewAutoload().JwtHelper()
	dToken, err := jwt.DecodeToken(token)

	/*
		validation token input
	*/
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message":    "invalid token",
		})
	}

	/*
		check token expiration
	*/
	expiredTime, err := dToken.GetExpirationTime()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message":    "error parse expired time",
		})
	}

	if time.Now().Unix() > expiredTime.Unix() {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message":    "token is expired",
		})
	}

	/*
		make sure user credential is valid
	*/
	userRepo := repositories.NewUser()
	userId, err := dToken.GetIssuer()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message":    "user not found",
		})
	}
	_, err = userRepo.FindById(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message":    "user not found",
		})
	}

	ctx.Next()
}
