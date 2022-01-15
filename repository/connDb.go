package repository

import (
	"github.com/go-redis/redis/v8"
)

type ConnectionRepository struct {
	Client *redis.Client
}

func (c *ConnectionRepository) InitConnectionRepository() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	c.Client = redisClient
}
