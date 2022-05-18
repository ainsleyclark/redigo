// Copyright 2020 The RediGo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redigo

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// Encoder defines methods for encoding and decoding
// buffers into the cache store.
type Encoder interface {
	Encode(value any) ([]byte, error)
	Decode([]byte, any) error
}

// GobEnc implements the encoder interface.
type GobEnc struct{}

func (g GobEnc) Encode(value any) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (g GobEnc) Decode(data []byte, value any) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(value)
}

// JSONEnc implements the encoder interface.
type JSONEnc struct{}

func (j JSONEnc) Encode(value any) ([]byte, error) {
	return json.Marshal(value)
}

func (j JSONEnc) Decode(data []byte, value any) error {
	return json.Unmarshal(data, value)
}
