package p1908gameofnim

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nimGame(t *testing.T) {
	for _, tc := range []struct {
		piles []int
		want  bool
	}{
		{[]int{1}, true},
		{[]int{1, 1}, false},
		{[]int{1, 2, 3}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.piles), func(t *testing.T) {
			require.Equal(t, tc.want, nimGame(tc.piles))
		})
	}
}

func nimGame(piles []int) bool {
	// Start with the base case of piles of 1
	// With [1], the first player wins
	// [1,1] => second player wins
	// [1,1,1] => first player wins, and so on.
	//
	// What about adding one to a single pile?
	// [2] => first player wins
	// [2,1] => second player wins
	// [2,2] => first player wins
	// [2,2,1] => second player wins (try all alternatives and it becomes clear)
	//
	// The pattern is: if XOR of the piles becomes zero, then the next player will
	// win. Any sum can also be adjusted to become non-zero.
	//
	sum := 0
	for _, p := range piles {
		sum ^= p
	}
	return sum > 0
}
