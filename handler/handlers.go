package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/haikalvidya/goShortURL/storage"
	"github.com/haikalvidya/goShortURL/shortener"
)

// request model init
type URLCreationReq struct{
	LongUrl string `json:"long_url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

func CreateShortURL(c *gin.Context){
	var createReq URLCreationReq
	if err := c.ShouldBindJSON(&createReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortUrl := shortener.GenerateShortLink(createReq.LongUrl, createReq.UserId)
	storage.SavedURLMapping(shortUrl, createReq.LongUrl, createReq.UserId)

	host := "http://localhost:9090/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host +shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := storage.RetrieveInitialURL(shortUrl)
	c.Redirect(302, initialUrl)
}