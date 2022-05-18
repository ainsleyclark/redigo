// Copyright 2020 The RediGo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redigo

import (
	"context"
	"github.com/ainsleyclark/redigo/internal"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type (
	// Cache defines the methods for interacting with the
	// cache layer.
	Cache struct {
		client  internal.RedisStore
		mtx     *sync.Mutex
		encoder Encoder
	}
	// Options represents the cache store available options
	// when using Set().
	Options struct {
		// Expiration allows to specify a global expiration
		// time hen setting a value.
		Expiration time.Duration
		// Tags allows specifying associated tags to the
		// current value.
		Tags []string
	}
	// Store defines methods for interacting with the
	// caching system.
	Store interface {
		// Ping pings the Redis cache to ensure its alive.
		Ping(context.Context) error
		// Get retrieves a specific item from the cache by key. Values are
		// automatically marshalled for use with Redis.
		Get(context.Context, string, any) error
		// Set stores a singular item in memory by key, value
		// and options (tags and expiration time). Values are automatically
		// marshalled for use with Redis & Memcache.
		Set(context.Context, string, any, Options) error
		// Delete removes a singular item from the cache by
		// a specific key.
		Delete(context.Context, string) error
		// Invalidate removes items from the cache via the tags passed.
		Invalidate(context.Context, []string)
		// Flush removes all items from the cache.
		Flush(context.Context)
	}
)

// New creates a new store to Redis instance(s).
func New(opts *redis.Options, enc Encoder) *Cache {
	return &Cache{
		client:  redis.NewClient(opts),
		mtx:     &sync.Mutex{},
		encoder: enc,
	}
}

// Ping pings the Redis cache to ensure its alive.
func (c *Cache) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

// Get retrieves a specific item from the cache by key. Values are
// automatically marshalled for use with Redis.
func (c *Cache) Get(ctx context.Context, key string, v any) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	result, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = c.encoder.Decode([]byte(result), v)
	if err != nil {
		return err
	}

	return nil
}

// Set stores a singular item in memory by key, value
// and options (tags and expiration time). Values are automatically
// marshalled for use with Redis & Memcache.
func (c *Cache) Set(ctx context.Context, key string, value any, options Options) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	buf, err := c.encoder.Encode(value)
	if err != nil {
		return err
	}

	err = c.client.Set(ctx, key, buf, options.Expiration).Err()
	if err != nil {
		return err
	}

	if len(options.Tags) > 0 {
		c.setTags(ctx, key, options.Tags)
	}

	return nil
}

//rm -rf mocks \
//&& mockery --all --keeptree --exported=true --output=./mocks \
//&& mv mocks/internal mocks/redis

// Delete removes a singular item from the cache by
// a specific key.
func (c *Cache) Delete(ctx context.Context, key string) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	_, err := c.client.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}

// Invalidate removes items from the cache from the tags passed.
func (c *Cache) Invalidate(ctx context.Context, tags []string) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if len(tags) == 0 {
		return
	}

	for _, tag := range tags {
		cacheKeys, err := c.client.SMembers(ctx, tag).Result()
		if err != nil {
			continue
		}

		for _, cacheKey := range cacheKeys {
			c.client.Del(ctx, cacheKey)
		}

		c.client.Del(ctx, tag)
	}
}

// Flush removes all items from the cache.
func (c *Cache) Flush(ctx context.Context) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.client.FlushAll(ctx)
}

// setTags sets SMembers in the redis store for caching.
func (c *Cache) setTags(ctx context.Context, key any, tags []string) {
	for _, tag := range tags {
		c.client.SAdd(ctx, tag, key.(string))
		c.client.Expire(ctx, tag, 720*time.Hour)
	}
}
