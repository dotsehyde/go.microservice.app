package cache

import (
	"context"
	"errors"
	"time"

	"github.com/BenMeredithConsult/locagri-apps/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri-apps/app/framework/database"
	rcache "github.com/go-redis/cache/v9"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type cache struct {
	client *rcache.Cache
}

func NewService(rdb *database.RedisAdapter) gateways.CacheService {
	mycache := rcache.New(
		&rcache.Options{
			Redis:      rdb.DB,
			LocalCache: rcache.NewTinyLFU(1000, time.Minute),
		},
	)
	return &cache{client: mycache}
}

func (c *cache) Set(key string, value any, ttl time.Duration) error {
	if err := c.client.Set(
		&rcache.Item{
			Ctx:   context.Background(),
			Key:   key,
			Value: value,
			TTL:   ttl,
		},
	); err != nil {
		return err
	}
	return nil
}
func (c *cache) Get(key string, obj any) error {
	if err := c.client.Get(context.Background(), key, obj); err != nil {
		return ErrKeyNotFound
	}
	return nil
}
func (c *cache) Has(key string) bool {
	return c.client.Exists(context.Background(), key)
}
func (c *cache) Delete(key string) error {
	return c.client.Delete(context.Background(), key)
}
