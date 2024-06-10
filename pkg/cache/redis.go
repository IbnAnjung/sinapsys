package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	c *redis.Client
}

func NewRedisCache(c *redis.Client) CacheService {
	return &redisCache{
		c,
	}
}

func (c *redisCache) getLockKey(key string) string {
	return fmt.Sprintf("lock_%s", key)
}

func (c *redisCache) Set(ctx context.Context, key string, val interface{}, expTime time.Duration) (err error) {
	s, _ := json.Marshal(val)
	fmt.Println("=> set cache", string(s))
	_, err = c.c.Set(ctx, key, string(s), expTime).Result()
	return
}

func (c *redisCache) Get(ctx context.Context, key string, val interface{}) (err error) {
	strVal, err := c.c.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return
	}

	fmt.Println("=>get", key, strVal)
	if strVal == "" {
		err = nil
		return
	}

	if err = json.Unmarshal([]byte(strVal), val); err != nil {
		return
	}

	return
}

func (c *redisCache) GetAndLock(ctx context.Context, key string, val interface{}) (err error) {
	i := 1
	for {
		isAvailable, err := c.c.SetNX(ctx, c.getLockKey(key), 1, 0).Result()
		if !isAvailable || err != nil {
			fmt.Println("not available")
			if i == 5 {
				return errors.New("error get value")
			}

			time.Sleep(time.Duration(i) * 300 * time.Millisecond)
			i++
		} else {
			fmt.Println("available")
			break
		}
	}

	return c.Get(ctx, key, val)
}

func (c *redisCache) Unlock(ctx context.Context, key string) (err error) {
	_, err = c.c.Del(ctx, c.getLockKey(key)).Result()
	return
}
