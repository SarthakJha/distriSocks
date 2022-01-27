package repository

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type ConnectionRepository struct {
	Client *redis.Client
}

func (c *ConnectionRepository) InitConnectionRepository(redisHost, redisPort string) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})

	c.Client = redisClient
}
