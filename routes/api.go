package routes

import "github.com/gin-gonic/gin"

func Api(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("/users/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
}
