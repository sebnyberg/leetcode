package p1680concatconseq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_concatenatedBinary(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		// {1, 1},
		// {3, 27},
		{12, 505379714},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, concatenatedBinary(tc.n))
		})
	}
}

func concatenatedBinary(n int) int {
	sum := 0
	var mod int = 1e9 + 7
	width := 0
	for i := 1; i <= n; i++ {
		if (i & (i - 1)) == 0 {
			width++
		}
		sum = ((sum << width) | i) % mod
	}
	return sum
}
