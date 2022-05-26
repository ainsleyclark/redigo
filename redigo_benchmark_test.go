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

var merges = []struct {
	name string
	enc  Encoder
}{
	{"JSON", NewJSONEncoder()},
	{"Gob", NewGobEncoder()},
	{"Message Pack", NewMessagePackEncoder()},
}

func BenchmarkEncode(b *testing.B) {
	for _, merge := range merges {
		b.ReportAllocs()
		m := createMap(100)
		b.Run(merge.name, func(b *testing.B) {
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				res, err := merge.enc.Encode(m)
				_ = res
				assert.NoError(b, err)
			}
			b.StopTimer()
		})
	}
}

func BenchmarkDecode(b *testing.B) {
	for _, merge := range merges {
		b.ReportAllocs()
		m := createMap(100)
		buf, err := merge.enc.Encode(m)
		assert.NoError(b, err)
		b.Run(merge.name, func(b *testing.B) {
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				m := make(map[int64]float64)
				err := merge.enc.Decode(buf, &m)
				assert.NoError(b, err)
			}
			b.StopTimer()
		})
	}
}

//func BenchmarkEncode_100(b *testing.B) {
//	for _, merge := range merges {
//		for k := 0.; k <= 10; k++ {
//			b.ReportAllocs()
//
//			n := int(math.Pow(2, k))
//			m := createMap(int(k))
//			b.Run(fmt.Sprintf("%s/%d", merge.name, n), func(b *testing.B) {
//				for i := 0; i < b.N; i++ {
//					b.StartTimer()
//					res, err := merge.enc.Encode(m)
//					_ = res
//					assert.NoError(b, err)
//				}
//				b.StopTimer()
//			})
//		}
//	}
//}
//
//func BenchmarkDecode_100(b *testing.B) {
//	for _, merge := range merges {
//		for k := 0.; k <= 10; k++ {
//			b.ReportAllocs()
//
//			n := int(math.Pow(2, k))
//			m := createMap(int(k))
//			buf, err := merge.enc.Encode(m)
//			assert.NoError(b, err)
//
//			b.Run(fmt.Sprintf("%s/%d", merge.name, n), func(b *testing.B) {
//				for i := 0; i < b.N; i++ {
//					b.StartTimer()
//					m := make(map[int64]float64)
//					err := merge.enc.Decode(buf, &m)
//					assert.NoError(b, err)
//				}
//				b.StopTimer()
//			})
//		}
//	}
//}
