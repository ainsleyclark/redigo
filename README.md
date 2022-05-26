<div align="center">
<img height="250" src="res/logo.svg" alt="Errors Logo" />

&nbsp;

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

### Message Pack
Use `NewMessagePackEncoder()` in the constructor when creating a new client.
See [github.com/vmihailenco/msgpack](https://github.com/vmihailenco/msgpack) for more details.

```go
c := redigo.New(&redis.Options{}, redigo.NewMessagePackEncoder())
```

### Custom
You can pass in custom encoders to the client constructor, that implement the Encode and Decode methods.

```go
type MyEncoder struct{}

func (m MessagePack) Encode(value any) ([]byte, error) {
	// Marshal or encode value
	return []byte("hello"), nil
}

func (m MessagePack) Decode(data []byte, value any) error {
	// Unmarshal or decode value
	return nil
}

func ExampleCustom() {
	c := redigo.New(&redis.Options{}, &MyEncoder{})
}
```


### Benchmarks

```bash
$ go version
go version go1.18.2 darwin/amd64

$ go test -benchmem -bench .
goos: darwin
goarch: amd64
pkg: github.com/ainsleyclark/redigo
cpu: AMD Ryzen 7 5800X 8-Core Processor
BenchmarkEncode/JSON-16                    56397             21239 ns/op            9293 B/op        206 allocs/op
BenchmarkEncode/Gob-16                    161018              7481 ns/op            4304 B/op        220 allocs/op
BenchmarkEncode/Message_Pack-16           110452             10688 ns/op            6819 B/op        208 allocs/op
BenchmarkDecode/JSON-16                    38413             31247 ns/op            7245 B/op        302 allocs/op
BenchmarkDecode/Gob-16                     58423             21294 ns/op           12732 B/op        193 allocs/op
BenchmarkDecode/Message_Pack-16            62198             19288 ns/op            7217 B/op        220 allocs/op
PASS
```

## Credits
Shout out to the incredible [Maria Letta](https://github.com/MariaLetta) for her excellent Gopher illustrations
