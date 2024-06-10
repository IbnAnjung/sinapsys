package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr           string
	Username       string
	Password       string
	ClientName     string
	DB             int
	MinIdleConns   int
	MaxIdleConns   int
	MaxActiveConns int
}

type Redis struct {
	Client *redis.Client
}

func NewRedis(ctx context.Context, config RedisConfig) (r Redis, err error) {
	rc := redis.NewClient(&redis.Options{
		Addr:           config.Addr,
		Username:       config.Username,
		Password:       config.Password,
		ClientName:     config.ClientName,
		DB:             config.DB,
		MinIdleConns:   config.MaxIdleConns,
		MaxIdleConns:   config.MaxIdleConns,
		MaxActiveConns: config.MaxActiveConns,
	})

	if err = rc.Ping(ctx).Err(); err != nil {
		return
	}
	r.Client = rc
	return
}

func (r *Redis) Cleanup() error {
	return r.Client.Close()
}
