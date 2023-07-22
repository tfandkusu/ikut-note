package main

import (
	"github.com/gin-gonic/gin"

	"fmt"
)

func main() {
	fmt.Println("Server Start")
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	err := router.Run(":4456")
	if err != nil {
		fmt.Println(err)
	}
}
