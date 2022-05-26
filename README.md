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
```

### Encode

```bash
BenchmarkEncode/JSON/Small-16            4942267               241.6 ns/op           424 B/op          3 allocs/op
BenchmarkEncode/JSON/Medium-16            776106              1550 ns/op            3484 B/op          3 allocs/op
BenchmarkEncode/JSON/Large-16              53864             22118 ns/op           42643 B/op          3 allocs/op
BenchmarkEncode/Gob/Small-16             1957748               626.6 ns/op          1096 B/op         13 allocs/op
BenchmarkEncode/Gob/Medium-16            1314542               903.4 ns/op          4360 B/op         13 allocs/op
BenchmarkEncode/Gob/Large-16              200398              6004 ns/op           58120 B/op         13 allocs/op
BenchmarkEncode/Message_Pack/Small-16    				 7899438               149.5 ns/op           424 B/op          4 allocs/op
BenchmarkEncode/Message_Pack/Medium-16           3438433               346.9 ns/op          2185 B/op          4 allocs/op
BenchmarkEncode/Message_Pack/Large-16             399356              2996 ns/op           28827 B/op          4 allocs/op
BenchmarkEncode/Go_JSON/Small-16                 6966218               151.6 ns/op           424 B/op          4 allocs/op
BenchmarkEncode/Go_JSON/Medium-16                3529380               347.9 ns/op          2185 B/op          4 allocs/op
BenchmarkEncode/Go_JSON/Large-16                  420511              2871 ns/op           28827 B/op          4 allocs/op
```

### Decode

```bash
BenchmarkDecode/JSON/Small-16                     348264              3517 ns/op            1665 B/op         47 allocs/op
BenchmarkDecode/JSON/Medium-16                     82147             14939 ns/op           10436 B/op        209 allocs/op
BenchmarkDecode/JSON/Large-16                       4554            262382 ns/op          210426 B/op       2794 allocs/op
BenchmarkDecode/Gob/Small-16                       79164             13980 ns/op            8911 B/op        226 allocs/op
BenchmarkDecode/Gob/Medium-16                      29278             41238 ns/op           24799 B/op        553 allocs/op
BenchmarkDecode/Gob/Large-16                        2779            423984 ns/op          222722 B/op       4660 allocs/op
BenchmarkDecode/Message_Pack/Small-16            1000000              1023 ns/op             938 B/op         25 allocs/op
BenchmarkDecode/Message_Pack/Medium-16            149588              8066 ns/op            8878 B/op        190 allocs/op
BenchmarkDecode/Message_Pack/Large-16               9164            128152 ns/op          138134 B/op       2493 allocs/op
BenchmarkDecode/Go_JSON/Small-16                 1210730               996.8 ns/op           938 B/op         25 allocs/op
BenchmarkDecode/Go_JSON/Medium-16                 148923              8087 ns/op            8861 B/op        190 allocs/op
BenchmarkDecode/Go_JSON/Large-16                    9086            128562 ns/op          138248 B/op       2493 allocs/op
```

## Credits
Shout out to the incredible [Maria Letta](https://github.com/MariaLetta) for her excellent Gopher illustrations
