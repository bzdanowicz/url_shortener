package storage

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Context context.Context
	Client  *redis.Client
}

func NewCache(context context.Context, address string) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})
	return &Cache{context, client}
}

func (c *Cache) Set(key, value string) error {
	return c.Client.Set(c.Context, key, value, 0).Err()
}

func (c *Cache) Get(key string) (string, error) {
	res, err := c.Client.Get(c.Context, key).Result()
	return res, err
}
