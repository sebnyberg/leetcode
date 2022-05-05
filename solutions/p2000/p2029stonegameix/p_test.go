package p2029stonegameix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_stoneGameIX(t *testing.T) {
	for _, tc := range []struct {
		stones []int
		want   bool
	}{
		{[]int{20, 3, 20, 17, 2, 12, 15, 17, 4}, true},
		// {[]int{2, 1}, true},
		// {[]int{2}, false},
		// {[]int{5, 1, 2, 4, 3}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.stones), func(t *testing.T) {
			require.Equal(t, tc.want, stoneGameIX(tc.stones))
		})
	}
}

func stoneGameIX(stones []int) bool {
	// If Alice picks 1 first, then Bob is forced to pick 1, then Alice is
	// forced to pick 2, then Bob 1, and so on.
	// [1,1,2,1,2,1,2,1]... until running out of 1s or 2s
	//
	// If Alice picks 2 first, then Bob is forced to pick 2, then Alice 1,
	// and so on:
	// [2,2,1,2,1,2,1,2]... until running out of 1s or 2s
	//
	// This means that Alice can guarantee the outcome, given that there are no
	// zeroes in the game.
	//
	// If there are an even number of zeroes, Alice can counter any zero with
	// another zero. Whether or not it is optimal doesn't matter.
	//
	// If there is an uneven number of zeroes, Bob can reverse Alice's initial
	// move, ensuring that he will always win.

	var count [3]int
	for _, stone := range stones {
		count[stone%3]++
	}
	if min(count[1], count[2]) == 0 {
		// If there is only one stone to pick that is non-zero,
		// Then Alice has to pick that one, then Bob picks it again next turn
		// On the third stone, Alice will lose, or if there is an odd number of
		// zeroes, then Alice can postpone picking the third stone until Bob is
		// forced to.
		return max(count[1], count[2]) > 2 && count[0]%2 == 1
	}
	// Alice will pick the most beneficial series
	// If count[0] is even, then Alice picks a number so that Bob ends up
	// with the wrong one.
	// If count[0] is odd, then Bob will be able to shift the series one step
	// in his favour, picking the opposite value of what Alice wants (+1)
	// Also, the first number picked will be picked twice (+1).
	return count[0]%2 == 0 || abs(count[1]-count[2]) > 2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
