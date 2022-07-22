package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v9"
)

// # Data in obj
type appCache struct {
	cache *redis.Client
}

// # Constructor
func NewAppCache(cache *redis.Client) AppCache {
	return appCache{cache: cache}
}

func (c appCache) Get(key string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	countryVal, err := c.cache.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, errors.New("cache: no documents in result")
	} else if err != nil {
		return nil, err
	}
	return &countryVal, nil
}

func (c appCache) Set(key string, data string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := c.cache.Set(ctx, key, data, 30*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c appCache) Clear() error {
	return nil
}
