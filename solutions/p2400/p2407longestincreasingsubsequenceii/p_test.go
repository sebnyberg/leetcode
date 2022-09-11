package p2407longestincreasingsubsequenceii

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLIS(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{4, 2, 1, 4, 3, 4, 5, 8, 15}, 3, 5},
		{[]int{7, 4, 5, 1, 8, 12, 4, 7}, 5, 4},
		{[]int{1, 5}, 1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, lengthOfLIS(tc.nums, tc.k))
		})
	}
}

func lengthOfLIS(nums []int, k int) int {
	// This is a shot in the dark, but I'll try to use a variant of Patience sort
	// where each pile is just a sorted list. If Go had something like a Tree,
	// then this would be trivial...
	stack := [][]int{}

	// When considering each number, find the first pile in the stack which
	// contains a number that the current number could be added to.
	// When a match is find, do an inorder insert of the number into that pile.
	// The number of piles in the stack is the answer to the question.

	// Also, insert sentinels to make code easier to manage
	for _, v := range nums {
		i := len(stack) // insert position
		for ; i-1 >= 0; i-- {
			j := sort.SearchInts(stack[i-1], v)
			if stack[i-1][j-1] >= v-k && stack[i-1][j-1] < v {
				break
			}
		}

		if i >= len(stack) {
			stack = append(stack, []int{math.MinInt32, v, math.MaxInt32})
		} else {
			j := sort.SearchInts(stack[i], v)
			if stack[i][j] != v {
				stack[i] = append(stack[i], 0)
				copy(stack[i][j+1:], stack[i][j:])
				stack[i][j] = v
			}
		}
	}
	return len(stack)
}
