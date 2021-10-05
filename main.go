package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bzdanowicz/url_shortener/storage"
	"github.com/bzdanowicz/url_shortener/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/bzdanowicz/url_shortener/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

type UrlCreationRequest struct {
	OriginalUrl string `json:"original_url" binding:"required"`
}

type UrlCreationResponse struct {
	Message     string `json:"message" binding:"required"`
	OriginalUrl string `json:"original_url" binding:"required"`
	NewUrl      string `json:"new_url" binding:"required"`
}

type UrlSearchResponse struct {
	Message     string `json:"message" binding:"required"`
	OriginalUrl string `json:"original_url" binding:"required"`
	NewUrl      string `json:"new_url" binding:"required"`
}

// FindOriginalUrl godoc
// @Summary Get original url.
// @Description returns the original page url
// @Accept json
// @Produce json
// @Param url path string true "url"
// @Success 202 "returns the original page url"
// @Router /{url} [get]
func FindOriginalUrl(c *gin.Context, storageManager *storage.StorageManager) {
	url := c.Param("url")
	originalUrl, err := storageManager.GetOriginalUrl(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, &UrlSearchResponse{"Link does not exist", "", ""})
		return
	}
	go storageManager.UpdateStatistics(url)
	c.JSON(http.StatusAccepted, &UrlSearchResponse{"Link exists", originalUrl, url})
}

// CreateNewLink godoc
// @Summary Creating new short link
// @Description shorten original link
// @Accept json
// @Produce json
// @Param req body UrlCreationRequest true "OriginalUrl"
// @Success 202 {object} UrlCreationResponse "Returns shorten url"
// @Failure 400 {object} UrlCreationResponse "Failure response"
// @Router /url [post]
func CreateNewLink(c *gin.Context, storageManager *storage.StorageManager) {
	request := &UrlCreationRequest{}
	err := c.Bind(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encodedUrl, err := utils.EncodeUrl(request.OriginalUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = storageManager.GetOriginalUrl(encodedUrl)
	if err == nil {
		c.JSON(http.StatusAccepted, &UrlCreationResponse{"Link already exists", request.OriginalUrl, encodedUrl})
		return
	}

	url, err := storageManager.CreateShortUrl(request.OriginalUrl, encodedUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, &UrlCreationResponse{"Created new link", request.OriginalUrl, url.Id})
}

func createRouting(r *gin.Engine, storageManager *storage.StorageManager) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome",
		})
	})

	r.GET("/:url", func(c *gin.Context) {
		FindOriginalUrl(c, storageManager)
	})

	r.POST("/url", func(c *gin.Context) {
		CreateNewLink(c, storageManager)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func runRouter(storageManager *storage.StorageManager) {
	router := gin.Default()
	router.Use(cors.Default())
	createRouting(router, storageManager)
	router.Run()
}

func main() {
	ctx := context.Background()
	manager := initializeApplication(ctx)
	runRouter(manager)
}
