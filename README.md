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

## Why?
RediGo is a wrapper for the Redis V8 GoLang client that features tagging, expiration and automatic encoding and decoding
using various encoders. It helps to unify various encoding techniques with a simple and easy to user interface.

Gob encoding performs drastically better in comparison to JSON which you can see from the benchmarks below.

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
	err = c.Get(ctx, "my-key", &val) // Be sure to pass a reference!
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

```go
c := redigo.New(&redis.Options{}, redigo.NewJSONEncoder())
```

### Gob

```go
c := redigo.New(&redis.Options{}, redigo.NewGobEncoder())
```

### Message Pack
See [github.com/vmihailenco/msgpack](https://github.com/vmihailenco/msgpack) for more details.

```go
c := redigo.New(&redis.Options{}, redigo.NewMessagePackEncoder())
```

### Go JSON
See [github.com/goccy/go-json](https://github.com/goccy/go-json) for more details.

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
```

### Encode

```bash
BenchmarkEncode/JSON-16                    54728             21813 ns/op            9294 B/op        206 allocs/op
BenchmarkEncode/Gob-16                    154272              7629 ns/op            4304 B/op        220 allocs/op
BenchmarkEncode/Message_Pack-16           113059             10468 ns/op            6820 B/op        208 allocs/op
BenchmarkEncode/Go_JSON-16                 92598             12768 ns/op             897 B/op          1 allocs/op
```

#### Graph representing ns/op.
<img width="100%" src="graph/Encode.svg" alt="Encoding Benchmark Graph" />

### Decode

```bash
BenchmarkDecode/JSON/-16                   39386             30318 ns/op            7246 B/op        302 allocs/op
BenchmarkDecode/Gob/-16                    57792             20742 ns/op           12733 B/op        193 allocs/op
BenchmarkDecode/Message_Pack/-16           57416             20626 ns/op            7217 B/op        220 allocs/op
BenchmarkDecode/Go_JSON/-16                95376             12186 ns/op            8068 B/op        220 allocs/op

```

#### Graph representing ns/op.
<img width="100%" src="graph/Decode.svg" alt="Decoding Benchmark Graph" />

## Contributing

Please feel free to make a pull request if you think something should be added to this package!

## Credits

Shout out to the incredible [Maria Letta](https://github.com/MariaLetta) for her excellent Gopher illustrations.

## Licence

Code Copyright 2022 RediGo. Code released under the [MIT Licence](LICENSE).
