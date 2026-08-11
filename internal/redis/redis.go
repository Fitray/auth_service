package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/Fitray/auth_service/internal/config"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	*redis.Client
	Timeout time.Duration
}

type Redis interface {
	Get(
		ctx context.Context,
		key string,
	) *redis.StringCmd

	Set(
		ctx context.Context,
		key string,
		value interface{},
		expiration time.Duration,
	) *redis.StatusCmd

	Del(
		ctx context.Context,
		keys ...string,
	) *redis.IntCmd

	Exists(
		ctx context.Context,
		keys ...string,
	) *redis.IntCmd
}

func NewClient(
	redisConfig config.RedisConfig,
) (*Client, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		redisConfig.Timeout,
	)
	defer cancel()

	client := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf(
				"%s:%s",
				redisConfig.Host, redisConfig.Port,
			),
			Password: redisConfig.Password,
			DB:       redisConfig.DB,
			PoolSize: redisConfig.PoolSize,
			Username: redisConfig.Username,
		},
	)

	if err := client.Ping(ctx).Err(); err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return &Client{
		Client:  client,
		Timeout: redisConfig.Timeout,
	}, nil
}

func (c Client) GetTimeout() time.Duration {
	return c.Timeout
}

func (c *Client) Close() {
	c.Client.Close()
}
