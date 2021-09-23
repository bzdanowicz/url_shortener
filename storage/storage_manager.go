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

func (manager *StorageManager) CreateShortURL(originalURL string) error {
	encodedURL, err := utils.EncodeURL(originalURL)
	if err != nil {
		return err
	}
	shortURL := &ShortURL{encodedURL, originalURL, 0}
	_, err = manager.Database.Insert(shortURL, "links")

	return err
}

func (manager *StorageManager) UpdateStatistics(shortURL string) error {
	document, err := manager.Database.FindShortURL(shortURL, "links")
	if err != nil {
		return fmt.Errorf("couldn't update statistics: %w", err)
	}
	_, err = manager.Database.UpdateVisits(document.Id, document.Visits+1, "links")
	return err
}

func (manager *StorageManager) GetOriginalURL(shortURL string) (string, error) {
	originalURL, err := manager.Cache.Get(shortURL)
	if err != nil {
		document, err := manager.Database.FindShortURL(shortURL, "links")
		if err != nil {
			return "", errors.New("couldn't find in cache and database")
		}
		return document.OriginalUrl, err
	}
	return originalURL, err
}
