package main

import (
	"fmt"
	"os"

	"github.com/christoffer1009/url-shortener-go/handler"
	"github.com/christoffer1009/url-shortener-go/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	PORT := os.Getenv("APP_PORT")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá, aqui é um encurtador de URLs em Go!",
		})
	})

	r.POST("/encurtar", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(fmt.Sprintf(":%v", PORT))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
