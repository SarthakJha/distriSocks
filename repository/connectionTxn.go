package repository

import (
	"context"
	"fmt"
	"time"
)

func (c *ConnectionRepository) GetWSConnections(recvID string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	val := c.Client.Get(ctx, recvID)
	valFinal, err := val.Result()
	if err != nil {
		return nil, err
	}
	return []string{valFinal}, nil
}

func (c *ConnectionRepository) SetWSConnection(key, val string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res := c.Client.Set(ctx, key, val, 0) // 0 means no expiration
	val, err := res.Result()
	fmt.Println(val)
	if err != nil {
		return err
	}
	return nil
}

func (c *ConnectionRepository) DeleteWSConnection(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result := c.Client.Del(ctx, key)
	_, err := result.Result()
	if err != nil {
		return nil
	}
	return err
}
