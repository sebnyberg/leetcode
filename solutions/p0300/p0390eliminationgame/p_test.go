package p0390eliminationgame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lastRemaining(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{9, 6},
		{4, 2},
		{1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, lastRemaining(tc.n))
		})
	}
}

func lastRemaining(n int) int {
	// Keep track of the first element, and the delta between numbers
	first := 1
	delta := 1

	nleft := n
	for even := true; ; even = !even {
		// For even operations, and odd operations where the number of elements
		// is odd, then the first element is replaced with the value first+delta.
		if even || nleft%2 == 1 {
			if first+delta > n {
				return first
			}
			first += delta
		}
		nleft /= 2
		delta <<= 1
	}
}
