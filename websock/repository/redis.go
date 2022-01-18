package repository

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

type ConnectionRepository struct {
	Client *redis.Client
}

func (c *ConnectionRepository) InitConnectionRepository() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})

	c.Client = redisClient
}
