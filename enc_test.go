package redigo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func UtilTestEncode(t *testing.T, enc Encoder, want string) {
	t.Helper()

	t.Run("Encode", func(t *testing.T) {
		tt := map[string]struct {
			input any
			err   bool
			want  any
		}{
			"Success": {
				"hello",
				false,
				want,
			},
			"Error": {
				make(chan []byte),
				true,
				nil,
			},
		}

		for name, test := range tt {
			t.Run(name, func(t *testing.T) {
				got, err := enc.Encode(test.input)
				if test.err {
					assert.Error(t, err)
					return
				}
				assert.Equal(t, test.want, string(got))
			})
		}
	})

	t.Run("Decode", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			val := ""
			err := enc.Decode([]byte(want), &val)
			assert.NoError(t, err)
		})
		t.Run("Error", func(t *testing.T) {
			val := ""
			err := enc.Decode([]byte("wrong"), &val)
			assert.Error(t, err)
		})
	})

}

func TestGobEncode(t *testing.T) {
	UtilTestEncode(t, NewGobEncoder(), "\b\f\x00\x05hello")
}

func TestJSONEncode(t *testing.T) {
	UtilTestEncode(t, NewJSONEncoder(), "\"hello\"")
}

func TestMessagePackEncode(t *testing.T) {
	UtilTestEncode(t, NewMessagePackEncoder(), "\xa5hello")
}
