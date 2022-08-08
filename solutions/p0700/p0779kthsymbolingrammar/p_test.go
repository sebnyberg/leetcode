package p0779kthsymbolingrammar

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthGrammar(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want int
	}{
		{3, 3, 1},
		{1, 1, 0},
		{2, 1, 0},
		{2, 2, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, kthGrammar(tc.n, tc.k))
		})
	}
}

func kthGrammar(n int, k int) int {
	var val int
	k--
	for k > 0 {
		w := 1 << int(math.Ceil(math.Log2(float64(k+1))))
		val ^= 1
		// Find k'th position in mirrored portion
		k -= w / 2
	}
	return val
}
