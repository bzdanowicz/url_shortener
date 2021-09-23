package main

import (
	"context"
	"log"

	"github.com/bzdanowicz/url_shortener/storage"
	"github.com/bzdanowicz/url_shortener/utils"
)

func initializeApplication(ctx context.Context) *storage.StorageManager {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Panicln("error during loading configuration, ", err)
	}

	database := storage.NewDatabase(ctx, config.DatabaseAddress, config.DatabaseName)
	cache := storage.NewCache(ctx, config.CacheAddress)
	return &storage.StorageManager{database, cache}
}

func main() {
	ctx := context.Background()
	initializeApplication(ctx)
}
