package p2350shortestimpossiblesequenceofrolls

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestSequence(t *testing.T) {
	for _, tc := range []struct {
		rolls []int
		k     int
		want  int
	}{
		{[]int{4, 2, 1, 2, 3, 3, 2, 4, 1}, 4, 3},
		{[]int{1, 1, 2, 2}, 2, 2},
		{[]int{1, 1, 3, 2, 2, 2, 3, 3}, 4, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rolls), func(t *testing.T) {
			require.Equal(t, tc.want, shortestSequence(tc.rolls, tc.k))
		})
	}
}

func shortestSequence(rolls []int, k int) int {
	seen := make([]byte, k+1)
	result := 1
	var bit byte = 1
	var count int
	for _, x := range rolls {
		if seen[x]&1 == bit {
			continue
		}
		seen[x] ^= 1
		count++
		if count == k {
			result++
			bit ^= 1
			count = 0
		}
	}
	return result
}
