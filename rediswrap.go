package rediswrap

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"time"
)

type Wrapper struct {
	redisClient redis.Client
}

func (w *Wrapper) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Redis Add")
	defer span.Finish()
	span.SetTag("key", key)
	span.SetTag("args", value)

	return w.redisClient.Set(ctx, key, value, expiration)
}

func (w *Wrapper) Get(ctx context.Context, key string) *redis.StringCmd {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Redis Get")
	defer span.Finish()
	span.SetTag("key", key)

	return w.redisClient.Get(ctx, key)
}

func (w *Wrapper) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Redis Del")
	defer span.Finish()
	span.SetTag("keys", keys)

	return w.redisClient.Del(ctx, keys...)
}