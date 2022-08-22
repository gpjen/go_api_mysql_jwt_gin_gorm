package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	group := router.Group("api/v1")

	group.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "mantep books",
		})
	})

	group.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "oke post",
		})
	})

	router.Run(":8080")
}
