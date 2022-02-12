package repository

import (
	"github.com/SarthakJha/distr-websock/internal/utils"
	"github.com/go-redis/redis/v8"
)

type ConnectionRepository struct {
	Client *redis.ClusterClient
}

func (c *ConnectionRepository) InitConnectionRepository(redisHost, redisPort string) {
	redisClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: utils.ResolveHeadlessServiceDNS(redisHost, "redis"),
	})

	c.Client = redisClient
}
