// Copyright 2020 The RediGo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type (
	// RedisStore is an abstraction of a *redis.Client used
	// for testing.
	RedisStore interface {
		Ping(ctx context.Context) *redis.StatusCmd
		Get(ctx context.Context, key string) *redis.StringCmd
		Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
		Del(ctx context.Context, keys ...string) *redis.IntCmd
		SMembers(ctx context.Context, key string) *redis.StringSliceCmd
		FlushAll(ctx context.Context) *redis.StatusCmd
		SAdd(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
		Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
		Close() error
	}
)
