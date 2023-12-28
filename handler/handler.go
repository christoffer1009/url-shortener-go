package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/christoffer1009/url-shortener-go/shortener"
	"github.com/christoffer1009/url-shortener-go/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type CreateShortUrlRequest struct {
	OriginalUrl string `json:"originalUrl" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var createUrlRequest CreateShortUrlRequest

	if err := c.ShouldBindJSON(&createUrlRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortUrl(createUrlRequest.OriginalUrl)
	store.SaveUrlMapping(shortUrl, createUrlRequest.OriginalUrl)

	godotenv.Load(".env")
	HOST := os.Getenv("APP_HOST")
	PORT := os.Getenv("APP_PORT")

	host := fmt.Sprintf("%v:%v/", HOST, PORT)
	c.JSON(200, gin.H{
		"shortUrl": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	originalUrl := store.GetOriginalUrl(shortUrl)
	c.Redirect(302, originalUrl)

}
