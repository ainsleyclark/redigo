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
			want  any
		}{
			"Success": {
				"hello",
				want,
			},
			"Error": {
				make(chan []byte),
				"type: chan []uint8",
			},
		}

		for name, test := range tt {
			t.Run(name, func(t *testing.T) {
				got, err := enc.Encode(test.input)
				if err != nil {
					assert.Contains(t, err.Error(), test.want)
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
	UtilTestEncode(t, &gobEnc{}, "\b\f\x00\x05hello")
}

func TestJSONEncode(t *testing.T) {
	UtilTestEncode(t, &jsonEnc{}, "\"hello\"")
}
