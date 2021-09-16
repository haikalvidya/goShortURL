package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haikalvidya/goShortURL/handler"
	"github.com/haikalvidya/goShortURL/storage"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey, welcome to GOShortURL API!",
		})
	})

	router.POST("/short-url", func(c *gin.Context){
		handler.CreateShortURL(c)
	})

	router.GET("/:shortURL", func(c *gin.Context){
		handler.HandleShortUrlRedirect(c)
	})

	// init storage
	storage.InitializeStorage()

	err := router.Run(":9090")
	if err != nil {
		panic(fmt.Sprintf("failed to start the web server - error: %v", err))
	}
}