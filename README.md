<div align="center">
<img height="250" src="res/logo.svg" alt="Errors Logo" />

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GoDoc](https://godoc.org/github.com/ainsleyclark/redigo/redis?status.svg)](https://pkg.go.dev/github.com/ainsleyclark/redigo)
[![Test](https://github.com/ainsleyclark/redigo/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/ainsleyclark/redigo/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/ainsleyclark/redigo/branch/master/graph/badge.svg?token=K27L8LS7DA)](https://codecov.io/gh/ainsleyclark/redigo)
[![GoReportCard](https://goreportcard.com/badge/github.com/ainsleyclark/redigo)](https://goreportcard.com/report/github.com/ainsleyclark/redigo)

</div>

# RediGo

A Redis client for GoLang featuring Tags with Gob &amp; JSON encoding.

## Install

```
go get -u github.com/ainsleyclark/redigo
```

## Quick Start

See below for a quick start to create a new Redis Client with an encoder. For more client methods see the
[Go Doc](https://pkg.go.dev/github.com/ainsleyclark/redigo) which includes all the client methods.

```go
func ExampleClient() {
	ctx := context.Background()

	c := redigo.New(&redis.Options{}, redigo.NewGobEncoder())
	err := c.Ping(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	err = c.Set(ctx, "my-key", "hello", redigo.Options{
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

```

## Encoders

### JSON
Use `NewJSONEncoder()` in the constructor when creating a new client.

```go
c := redigo.New(&redis.Options{}, redigo.NewJSONEncoder())
```

### Gob
Use `NewGobEncoder()` in the constructor when creating a new client.

```go
c := redigo.New(&redis.Options{}, redigo.NewGobEncoder())
```

### Custom
You can pass in custom encoders to the client constructor. Below is a message pack example. Using
[github.com/vmihailenco/msgpack](https://github.com/vmihailenco/msgpack)

```go
import "github.com/vmihailenco/msgpack/v5"

type MessagePack struct{}

func (m MessagePack) Encode(value any) ([]byte, error) {
	return msgpack.Marshal(value)
}

func (m MessagePack) Decode(data []byte, value any) error {
	return msgpack.Unmarshal(data, value)
}

func ExampleMessagePack() {
	c := redigo.New(&redis.Options{}, &MessagePack{})
}
```

## TODO
- Benchmarks

## Credits
Shout out to the incredible [Maria Letta](https://github.com/MariaLetta) for her excellent Gopher illustrations
