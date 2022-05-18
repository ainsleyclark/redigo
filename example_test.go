// Copyright 2020 The RediGo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redigo_test

import (
	"context"
	"github.com/ainsleyclark/redigo"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func Example() {
	ctx := context.Background()

	c := redigo.New(&redis.Options{}, redigo.NewGobEncoder())
	err := c.Ping(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	err := c.Set(ctx, "my-key", "hello", redigo.Options{
		Expiration: time.Second * 100,
		Tags:       []string{"my-tag"},
	})
	if err != nil {
		log.Fatalln(err)
	}

	var val string
	err = c.Get(ctx, "my-key", &val)
	if err != nil {
		log.Fatalln(err)
	}

	err = c.Delete(ctx, "my-key")
	if err != nil {
		log.Fatalln(err)
	}
}
