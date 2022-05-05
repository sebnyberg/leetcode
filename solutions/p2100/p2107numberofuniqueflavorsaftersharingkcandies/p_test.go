package p2107numberofuniqueflavorsaftersharingkcandies

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shareCandies(t *testing.T) {
	for _, tc := range []struct {
		candies []int
		k       int
		want    int
	}{
		{[]int{1, 2, 2, 3, 4, 3}, 3, 3},
		{[]int{2, 2, 2, 2, 3, 3}, 2, 2},
		{[]int{2, 4, 5}, 0, 3},
		{[]int{2, 4, 5}, 3, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.candies), func(t *testing.T) {
			require.Equal(t, tc.want, shareCandies(tc.candies, tc.k))
		})
	}
}

func shareCandies(candies []int, k int) int {
	var candyCount [1e5 + 1]uint16
	// 1. Parse all candies in the list, counting number of unique ones
	var nunique int
	for _, candy := range candies {
		if candyCount[candy] == 0 {
			nunique++
		}
		candyCount[candy]++
	}

	// 2. Move a window from left to right, counting maximum number of unique
	// candies
	var maxUnique int
	for i := 0; i < len(candies); i++ {
		candyCount[candies[i]]--
		if candyCount[candies[i]] == 0 {
			nunique--
		}
		if i >= k {
			candyCount[candies[i-k]]++
			if candyCount[candies[i-k]] == 1 {
				nunique++
			}
		}
		if i >= k-1 {
			maxUnique = max(maxUnique, nunique)
		}
	}
	return maxUnique
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
