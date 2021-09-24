package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bzdanowicz/url_shortener/storage"
	"github.com/bzdanowicz/url_shortener/utils"
	"github.com/gin-gonic/gin"
)

func initializeApplication(ctx context.Context) *storage.StorageManager {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Panicln("error during loading configuration, ", err)
	}

	database := storage.NewDatabase(ctx, config.DatabaseAddress, config.DatabaseName)
	cache := storage.NewCache(ctx, config.CacheAddress)
	return &storage.StorageManager{Database: database, Cache: cache}
}

func HandleRedirect(c *gin.Context, storageManager *storage.StorageManager) {
	url := c.Param("url")
	originalUrl, err := storageManager.GetOriginalUrl(url)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
	}
	go storageManager.UpdateStatistics(url)
	c.Redirect(http.StatusFound, originalUrl)
}

type UrlCreationRequest struct {
	OriginalUrl string `json:"original_url" binding:"required"`
}

func CreateNewLink(c *gin.Context, storageManager *storage.StorageManager) {
	request := &UrlCreationRequest{}
	err := c.Bind(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	url, err := storageManager.CreateShortUrl(request.OriginalUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Created new link", "original_url": request.OriginalUrl, "new_url": url.Id})
}

func createRouting(r *gin.Engine, storageManager *storage.StorageManager) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome",
		})
	})

	r.GET("/:url", func(c *gin.Context) {
		HandleRedirect(c, storageManager)
	})

	r.POST("/url", func(c *gin.Context) {
		CreateNewLink(c, storageManager)
	})
}

func runRouter(storageManager *storage.StorageManager) {
	router := gin.Default()
	createRouting(router, storageManager)
	router.Run()
}

func main() {
	ctx := context.Background()
	manager := initializeApplication(ctx)
	runRouter(manager)
}
