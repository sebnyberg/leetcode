package p1959minimumtotalspacewastedwithkresizingoptions

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSpaceWastedKResizing(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{10, 20}, 0, 10},
		{[]int{10, 20, 30}, 1, 10},
		{[]int{10, 20, 15, 30, 20}, 2, 15},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minSpaceWastedKResizing(tc.nums, tc.k))
		})
	}
}

func minSpaceWastedKResizing(nums []int, k int) int {
	n := len(nums)
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, k+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	return findMinSpace(mem, nums, k, 0, n)
}

func findMinSpace(mem [][]int, nums []int, k, i, n int) int {
	if i == n {
		return 0
	}
	if k == -1 {
		return math.MaxInt32
	}
	if mem[i][k] != -1 {
		return mem[i][k]
	}
	var maxNum, sum int
	res := math.MaxInt32
	// Evaluate resizing at position j somewhere between i <= j < n
	// Resizing must be done in accordance with the maximum element in the range
	for j := i; j < n; j++ {
		maxNum = max(maxNum, nums[j])
		sum += nums[j]
		wastedSpace := maxNum*(j-i+1) - sum
		res = min(res, wastedSpace+findMinSpace(mem, nums, k-1, j+1, n))
	}
	mem[i][k] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
