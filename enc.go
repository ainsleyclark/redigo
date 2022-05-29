// Copyright 2020 The RediGo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redigo

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	gojson "github.com/goccy/go-json"
	"github.com/vmihailenco/msgpack/v5"
)

// Encoder defines methods for encoding and decoding
// buffers into the cache store.
type Encoder interface {
	Encode(value any) ([]byte, error)
	Decode([]byte, any) error
}

// NewGobEncoder returns a new Gob encoder for RediGo.
func NewGobEncoder() Encoder {
	return &gobEnc{}
}

// gobEnc implements the encoder interface.
type gobEnc struct{}

func (g gobEnc) Encode(value any) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (g gobEnc) Decode(data []byte, value any) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(value)
}

// NewJSONEncoder returns a new Gob encoder for RediGo.
func NewJSONEncoder() Encoder {
	return &jsonEnc{}
}

// jsonEnc implements the encoder interface.
type jsonEnc struct{}

func (j jsonEnc) Encode(value any) ([]byte, error) {
	return json.Marshal(value)
}

func (j jsonEnc) Decode(data []byte, value any) error {
	return json.Unmarshal(data, value)
}

// NewMessagePackEncoder returns a new Message Pack
// encoder for RediGo.
func NewMessagePackEncoder() Encoder {
	return &msgEnc{}
}

// msgEnc implements the encoder interface.
type msgEnc struct{}

func (m msgEnc) Encode(value any) ([]byte, error) {
	return msgpack.Marshal(value)
}

func (m msgEnc) Decode(data []byte, value any) error {
	return msgpack.Unmarshal(data, value)
}

// NewGoJSONEncoder returns a new Go JSON
// encoder for RediGo.
func NewGoJSONEncoder() Encoder {
	return &goJSONEnc{}
}

// goJSONEnc implements the encoder interface.
type goJSONEnc struct{}

func (g goJSONEnc) Encode(value any) ([]byte, error) {
	return gojson.Marshal(value)
}

func (g goJSONEnc) Decode(data []byte, value any) error {
	return gojson.Unmarshal(data, value)
}
