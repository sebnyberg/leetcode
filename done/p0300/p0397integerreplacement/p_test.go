package p0397integerreplacement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_integerReplacement(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{8, 3},
		{7, 4},
		{4, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, integerReplacement(tc.n))
		})
	}
}

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	// Brute-force (BFS)
	// Normally I'd go through the different cases and check which is optimal.
	cur := map[int]struct{}{n: {}}
	moves := 1
	for {
		next := make(map[int]struct{}, len(cur))
		for alt := range cur {
			if alt%2 == 0 {
				next[alt/2] = struct{}{}
			} else {
				next[alt-1] = struct{}{}
				next[alt+1] = struct{}{}
			}
		}
		if _, exists := next[1]; exists {
			return moves
		}
		cur = next
		moves++
	}
}
