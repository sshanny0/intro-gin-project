package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HELLO WORLD!",
		})
	})

	router.Run(":8080")
}
