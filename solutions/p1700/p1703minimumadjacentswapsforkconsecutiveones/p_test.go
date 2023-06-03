package p1703minimumadjacentswapsforkconsecutiveones

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minMoves(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 0, 0, 0, 0, 0, 1, 1}, 3, 5},
		{[]int{1, 0, 0, 1, 0, 1}, 2, 1},
		{[]int{1, 1, 0, 1}, 2, 0},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minMoves(tc.nums, tc.k))
		})
	}
}

func minMoves(nums []int, k int) int {
	// The minimum cost of moving a group of 1s is the minimum cost of gathering
	// around the median of a value.
	// If k % 2 == 1, then we can pick each 1 and find the cost by combining
	// distance on left/right side to the next k/2 numbers.
	// If k % 2 == 0, then we do the same, but must pick a "preferred" side.

	// right[i] = cost of having l consecutive 1s to the right of i
	right := costToGatherLeft(nums, k/2)
	rev(right)
	rev(nums)
	left := costToGatherLeft(nums, k/2-(1-k&1))
	res := math.MaxInt32
	for i := range nums {
		if nums[i] == 0 {
			continue
		}
		res = min(res, left[i]+right[i])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func rev(nums []int) {
	n := len(nums)
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}

func costToGatherLeft(nums []int, k int) []int {
	n := len(nums)
	left := make([]int, n)
	if k == 0 {
		return left
	}
	for i := range left {
		left[i] = math.MaxInt32
	}

	queue := []int{}
	var dist int
	for i := range nums {
		if len(queue) == k {
			left[i] = dist
		}
		if nums[i] == 0 {
			// Move numbers one step forward
			dist += len(queue)
		} else {
			if len(queue) == k {
				// Pop oldest element
				dist -= (i - k) - queue[0]
				queue = queue[1:]
			}
			queue = append(queue, i)
		}
	}
	return left
}
