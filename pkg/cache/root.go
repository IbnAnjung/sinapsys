package cache

import (
	"context"
	"time"
)

type CacheService interface {
	Set(ctx context.Context, key string, val interface{}, expTime time.Duration) error
	Get(ctx context.Context, key string, val interface{}) (err error)
	GetAndLock(ctx context.Context, key string, val interface{}) (err error)
	Unlock(ctx context.Context, key string) (err error)
}
