package database

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"

	"github.com/BenMeredithConsult/locagri-apps/config"
)

type RedisAdapter struct {
	DB *redis.Client
}

func RedisDB() *RedisAdapter {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:6379", config.App().Redis),
			Password: "", // no password set
			DB:       0,  // use default DB
		},
	)
	_, e := rdb.Ping(context.Background()).Result()
	if e != nil {
		log.Panicln(e)
	}
	return &RedisAdapter{DB: rdb}
}
