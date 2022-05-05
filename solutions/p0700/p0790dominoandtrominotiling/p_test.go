package p0790dominoandtrominotiling

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numTilings(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{4, 11},
		{3, 5},
		{1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, numTilings(tc.n))
		})
	}
}

const mod = 1e9 + 7

func numTilings(n int) int {
	// Keep track of 2x2 boards
	// Lay down whatever domino/tromino's that are necessary to keep the first
	// column filled during the next round.
	var cur [1 << 2]int

	const (
		topleft = 1 << 0
		botleft = 1 << 1
	)
	cur[0] = 1

	for i := 0; i < n; i++ {
		var next [1 << 2]int

		// Given nothing:
		next[0] += cur[0]
		next[topleft] += cur[0]         // tri
		next[botleft] += cur[0]         // tri
		next[topleft|botleft] += cur[0] // two horizontal

		// Given topleft
		next[botleft] += cur[topleft]         // horizontal
		next[topleft|botleft] += cur[topleft] // tri

		// Given botleft
		next[topleft] += cur[botleft]         // horizontal
		next[topleft|botleft] += cur[botleft] // tri

		// Given both
		next[0] += cur[topleft|botleft]

		next[0] %= mod
		next[topleft] %= mod
		next[botleft] %= mod
		next[topleft|botleft] %= mod

		// switch
		cur = next
	}
	return cur[0] % mod
}
