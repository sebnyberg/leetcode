package p0393utf8validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validUtf8(t *testing.T) {
	for _, tc := range []struct {
		data []int
		want bool
	}{
		{[]int{237}, false},
		{[]int{197, 130, 1}, true},
		{[]int{235, 140, 4}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.data), func(t *testing.T) {
			require.Equal(t, tc.want, validUtf8(tc.data))
		})
	}
}

func validUtf8(data []int) bool {
	var remains int // remaining bytes in current code point
	for i := 0; i < len(data); i++ {
		b := byte(data[i])
		switch {
		case remains > 0:
			if b&0b11000000 != 0b10000000 {
				return false
			}
			remains--
		case b&0b10000000 == 0:
		case b&0b11100000 == 0b11000000:
			remains = 1
		case b&0b11110000 == 0b11100000:
			remains = 2
		case b&0b11111000 == 0b11110000:
			remains = 3
		default:
			return false
		}
	}
	return remains == 0
}
