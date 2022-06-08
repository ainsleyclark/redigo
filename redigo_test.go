// Copyright 2020 The RediGo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redigo

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/ainsleyclark/redigo/mocks"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
	"time"
)

// CacheTestSuite defines the helper used for
// cache testing.
type CacheTestSuite struct {
	suite.Suite
	GobBuf []byte
}

// TestCache asserts testing has begun.
func TestCache(t *testing.T) {
	suite.Run(t, &CacheTestSuite{})
}

// SetupSuite marshals bytes of the test value and assigns.
func (t *CacheTestSuite) SetupSuite() {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		t.Fail(err.Error())
	}
	t.GobBuf = buf.Bytes()
}

// Setup is a helper to obtain a mock cache store for testing.
func (t *CacheTestSuite) Setup(mf func(m *mocks.RedisStore, enc *mocks.Encoder)) *Cache {
	m := &mocks.RedisStore{}
	e := &mocks.Encoder{}
	if mf != nil {
		mf(m, e)
	}
	return &Cache{
		client:  m,
		mtx:     &sync.Mutex{},
		encoder: e,
	}
}

type (
	// testCacheStruct represents a struct for working with
	// JSON values within the cache store.
	testCacheStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}
)

var (
	// key is the test key used for Redis testing.
	key = "key"
	// tag is the test tag used for Redis testing.
	tag = "tag"
	// value is the test value to match against testing for
	// get and set test methods, to see if it's marshalling
	// properly.
	value = testCacheStruct{
		Name:  "name",
		Value: 1,
	}
	// options are the default testing set options.
	options = Options{
		Expiration: -1,
		Tags:       []string{tag},
	}
	ctx = context.TODO()
)

func (t *CacheTestSuite) TestNew() {
	got := New(&redis.Options{}, NewGobEncoder())
	t.NotNil(got.client)
	t.NotNil(got.mtx)
	t.NotNil(got.encoder)
}

func (t *CacheTestSuite) TestPing() {
	tt := map[string]struct {
		mock func(m *mocks.RedisStore, enc *mocks.Encoder)
		want any
	}{
		"Success": {
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				m.On("Ping", ctx).
					Return(redis.NewStatusCmd(ctx, nil))
			},
			nil,
		},
		"Ping Error": {
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				cmd := redis.NewStatusCmd(ctx, nil)
				cmd.SetErr(errors.New("ping error"))
				m.On("Ping", ctx).
					Return(cmd)
			},
			"ping error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			err := c.Ping(ctx)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
		})
	}
}

func (t *CacheTestSuite) TestClose() {
	tt := map[string]struct {
		mock func(m *mocks.RedisStore, enc *mocks.Encoder)
		want any
	}{
		"Success": {
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				m.On("Close").
					Return(nil)
			},
			nil,
		},
		"Ping Error": {
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				m.On("Close").
					Return(errors.New("close error"))
			},
			"close error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			err := c.Close()
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
		})
	}
}

func (t *CacheTestSuite) TestCache_Get() {
	tt := map[string]struct {
		mock func(m *mocks.RedisStore, enc *mocks.Encoder)
		want any
	}{
		"Success": {
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				m.On("Get", mock.Anything, key).
					Return(redis.NewStringResult(string(t.GobBuf), nil))

				enc.On("Decode", t.GobBuf, &testCacheStruct{}).
					Return(nil).
					Run(func(args mock.Arguments) {
						arg := args.Get(1).(*testCacheStruct)
						*arg = value
					})
			},
			value,
		},
		"Redis Error": {
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				m.On("Get", mock.Anything, key).
					Return(redis.NewStringResult("", fmt.Errorf("redis error")))
			},
			"redis error",
		},
		"Decode Error": {
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				m.On("Get", mock.Anything, key).
					Return(redis.NewStringResult(string(t.GobBuf), nil))

				enc.On("Decode", t.GobBuf, &testCacheStruct{}).
					Return(fmt.Errorf("decode error"))
			},
			"decode error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			got := testCacheStruct{}
			err := c.Get(ctx, key, &got)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
			t.Equal(nil, err)
		})
	}
}

func (t *CacheTestSuite) TestCache_Set() {
	tt := map[string]struct {
		value  any
		mock   func(m *mocks.RedisStore, enc *mocks.Encoder)
		panics bool
		want   any
	}{
		"Success": {
			value,
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				enc.On("Encode", value).
					Return(t.GobBuf, nil)

				m.On("Set", mock.Anything, key, t.GobBuf, options.Expiration).
					Return(redis.NewStatusCmd(ctx, nil))

				m.On("SAdd", ctx, "tag", "key").
					Return(redis.NewIntCmd(ctx, ""))

				m.On("Expire", ctx, "tag", 720*time.Hour).
					Return(redis.NewBoolCmd(ctx, true))
			},
			false,
			nil,
		},
		"Redis Error": {
			value,
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				enc.On("Encode", value).
					Return(t.GobBuf, nil)

				cmd := redis.NewStatusCmd(ctx, nil)
				cmd.SetErr(errors.New("redis error"))

				m.On("Set", mock.Anything, key, t.GobBuf, options.Expiration).
					Return(cmd)
			},
			true,
			"redis error",
		},
		"Encode Error": {
			value,
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				enc.On("Encode", value).
					Return(nil, fmt.Errorf("encode error"))
			},
			true,
			"encode error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			err := c.Set(ctx, key, test.value, options)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
		})
	}
}

func (t *CacheTestSuite) TestCache_Delete() {
	tt := map[string]struct {
		value any
		mock  func(m *mocks.RedisStore, enc *mocks.Encoder)
		error bool
		want  any
	}{
		"Success": {
			value,
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				m.On("Del", mock.Anything, key).
					Return(redis.NewIntCmd(ctx, nil))
			},
			false,
			nil,
		},
		"Redis Error": {
			value,
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				cmd := redis.NewIntCmd(ctx, nil)
				cmd.SetErr(errors.New("delete error"))

				m.On("Del", mock.Anything, key).
					Return(cmd)
			},
			true,
			"delete error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			err := c.Delete(ctx, key)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
		})
	}
}

func (t *CacheTestSuite) TestCache_Invalidate() {
	tt := map[string]struct {
		input []string
		mock  func(m *mocks.RedisStore, enc *mocks.Encoder)
	}{
		"Success": {
			[]string{tag},
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				cmd := redis.NewStringSliceCmd(ctx)
				cmd.SetVal([]string{key})
				m.On("SMembers", ctx, "tag").
					Return(cmd)

				m.On("Del", ctx, key).
					Return(redis.NewIntCmd(ctx, nil)).Once()

				m.On("Del", ctx, "tag").
					Return(redis.NewIntCmd(ctx, nil)).Once()
			},
		},
		"Nil Tags": {
			nil,
			nil,
		},
		"SMembers Error": {
			[]string{tag},
			func(m *mocks.RedisStore, enc *mocks.Encoder) {
				cmd := redis.NewStringSliceCmd(ctx)
				cmd.SetErr(errors.New("err"))

				m.On("SMembers", ctx, "tag").
					Return(cmd)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			c.Invalidate(ctx, test.input)
		})
	}
}

func (t *CacheTestSuite) TestCache_Flush() {
	c := t.Setup(func(m *mocks.RedisStore, enc *mocks.Encoder) {
		m.On("FlushAll", ctx).
			Return(redis.NewStatusCmd(ctx, nil))
	})
	c.Flush(ctx)
}
