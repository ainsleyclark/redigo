// Copyright 2020 The RediGo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redigo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	merges = []struct {
		name string
		enc  Encoder
	}{
		{"JSON", NewJSONEncoder()},
		{"Gob", NewGobEncoder()},
		{"Message Pack", NewMessagePackEncoder()},
		{"Go JSON", NewGoJSONEncoder()},
	}
)

func createMap(max int) map[int64]float64 {
	m := make(map[int64]float64)
	for i := 0; i < max; i++ {
		m[int64(i)] = float64(i)
	}
	return m
}

func BenchmarkEncode(b *testing.B) {
	b.ReportAllocs()

	for _, merge := range merges {
		b.Run(merge.name, func(b *testing.B) {
			m := createMap(100)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				res, err := merge.enc.Encode(m)
				_ = res
				if err != nil {
					b.Logf("Error during benchmark: %s", err.Error())
				}
			}
		})

	}
}

func BenchmarkDecode(b *testing.B) {
	b.ReportAllocs()

	for _, merge := range merges {
		b.Run(merge.name+"/", func(b *testing.B) {
			buf, err := merge.enc.Encode(createMap(100))
			assert.NoError(b, err)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m := make(map[int]float64)
				err := merge.enc.Decode(buf, &m)
				if err != nil {
					b.Logf("Error during benchmark: %s", err.Error())
				}
			}
		})
	}
}
