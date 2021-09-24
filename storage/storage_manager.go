package storage

import (
	"errors"
	"fmt"

	"github.com/bzdanowicz/url_shortener/utils"
)

type StorageManager struct {
	Database *Database
	Cache    *Cache
}

func (manager *StorageManager) CreateShortUrl(originalUrl string) (*ShortUrl, error) {
	encodedUrl, err := utils.EncodeUrl(originalUrl)
	if err != nil {
		return nil, err
	}
	shortUrl := &ShortUrl{encodedUrl, originalUrl, 0}
	_, err = manager.Database.Insert(shortUrl, "links")
	if err != nil {
		return nil, err
	}
	return shortUrl, err
}

func (manager *StorageManager) UpdateStatistics(shortUrl string) error {
	document, err := manager.Database.FindShortUrl(shortUrl, "links")
	if err != nil {
		return fmt.Errorf("couldn't update statistics: %w", err)
	}
	_, err = manager.Database.UpdateVisits(document.Id, document.Visits+1, "links")
	return err
}

func (manager *StorageManager) GetOriginalUrl(shortUrl string) (string, error) {
	originalUrl, err := manager.Cache.Get(shortUrl)
	if err != nil {
		document, err := manager.Database.FindShortUrl(shortUrl, "links")
		if err != nil {
			return "", errors.New("couldn't find in cache and database")
		}
		return document.OriginalUrl, err
	}
	return originalUrl, err
}
