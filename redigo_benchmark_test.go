// Copyright 2020 The RediGo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redigo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createMap(max int) map[int64]float64 {
	m := make(map[int64]float64)
	for i := 0; i < max; i++ {
		m[int64(i)] = float64(i)
	}
	return m
}

func BenchmarkEncode_Gob(b *testing.B) {
	b.ReportAllocs()
	m := createMap(1000)
	enc := NewGobEncoder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := enc.Encode(m)
		_ = res
		assert.NoError(b, err)
	}
}

func BenchmarkEncode_JSON(b *testing.B) {
	b.ReportAllocs()
	m := createMap(1000)
	enc := NewJSONEncoder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := enc.Encode(m)
		_ = res
		assert.NoError(b, err)
	}
}

func BenchmarkEncode_MessagePack(b *testing.B) {
	b.ReportAllocs()
	m := createMap(1000)
	enc := NewMessagePackEncoder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := enc.Encode(m)
		_ = res
		assert.NoError(b, err)
	}
}

func BenchmarkDecode_Gob(b *testing.B) {
	b.ReportAllocs()
	m := createMap(1000)
	enc := NewGobEncoder()
	res, err := enc.Encode(m)
	assert.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result map[int64]float64
		err := enc.Decode(res, &result)
		assert.NoError(b, err)
	}
}

func BenchmarkDecode_JSON(b *testing.B) {
	b.ReportAllocs()
	m := createMap(1000)
	enc := NewJSONEncoder()
	res, err := enc.Encode(m)
	assert.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result map[int64]float64
		err := enc.Decode(res, &result)
		assert.NoError(b, err)
	}
}

func BenchmarkDecode_MessagePack(b *testing.B) {
	b.ReportAllocs()
	m := createMap(1000)
	enc := NewMessagePackEncoder()
	res, err := enc.Encode(m)
	assert.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result map[int64]float64
		err := enc.Decode(res, &result)
		assert.NoError(b, err)
	}
}
