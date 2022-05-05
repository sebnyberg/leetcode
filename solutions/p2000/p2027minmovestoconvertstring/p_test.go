package p2027minmovestoconvertstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumMoves(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"XXX", 1},
		{"XX0X", 2},
		{"0000", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, minimumMoves(tc.s))
		})
	}
}

func minimumMoves(s string) int {
	var moves int
	for i := 0; i < len(s); {
		if s[i] == 'X' {
			moves++
			i += 3
		} else {
			i++
		}
	}
	return moves
}
