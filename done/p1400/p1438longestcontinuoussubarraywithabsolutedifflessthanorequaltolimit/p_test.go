package p1438longestcontinuoussubarraywithabsolutedifflessthanorequaltolimit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		limit int
		want  int
	}{
		{[]int{8, 2, 4, 7}, 4, 2},
		{[]int{10, 1, 2, 4, 7, 2}, 5, 4},
		{[]int{4, 2, 2, 2, 4, 4, 2, 2}, 0, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, longestSubarray(tc.nums, tc.limit))
		})
	}
}

func longestSubarray(nums []int, limit int) int {
	// Use two deques to keep a list of increasing / decreasing numbers by index
	minQ := make([]int, 0)
	maxQ := make([]int, 0)
	var i int
	for _, num := range nums {
		// Update deques
		for len(minQ) > 0 && num < minQ[len(minQ)-1] {
			minQ = minQ[:len(minQ)-1]
		}
		for len(maxQ) > 0 && num > maxQ[len(maxQ)-1] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		minQ = append(minQ, num)
		maxQ = append(maxQ, num)
		if maxQ[0]-minQ[0] > limit {
			// Pop oldest element at position i
			if maxQ[0] == nums[i] {
				maxQ = maxQ[1:]
			}
			if minQ[0] == nums[i] {
				minQ = minQ[1:]
			}
			i++
		}
	}
	return len(nums) - i
}
