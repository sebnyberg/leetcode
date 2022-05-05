package p1574shortestsubarraytoberemovedtomakearraysorted

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLengthOfShortestSubarray(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{1, 2, 2, 2, 2, 2, 3, 1, 7, 5, 1, 2, 2, 2, 2, 2, 2, 5, 6}, 5},
		{[]int{2, 2, 2, 1, 1, 1}, 3},
		{[]int{13, 0, 14, 7, 18, 18, 18, 16, 8, 15, 20}, 8},
		{[]int{1, 2, 3, 10, 0, 7, 8, 9}, 2},
		{[]int{1, 2, 3, 10, 4, 2, 3, 5}, 3},
		{[]int{1, 2, 3, 4, 3}, 1},
		{[]int{5, 4, 1, 2, 3, 4}, 2},
		{[]int{5, 4, 3, 2, 1}, 4},
		{[]int{1, 2, 3}, 0},
		{[]int{1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, findLengthOfShortestSubarray(tc.arr))
		})
	}
}

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)

	// Find non-decreasing prefix
	var left int
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			left = i
			break
		}
	}
	if left == 0 { // array is already non-decreasing
		return 0
	}

	// Find non-decreasing suffix
	stack := make([]int, 1, len(arr))
	stack[0] = left
	for i := left + 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			stack = stack[:1]
			stack[0] = i
			continue
		}
		stack = append(stack, i)
	}

	// Find optimal removal based on two non-decreasing subarrays
	res := len(arr)
	for i := 0; len(stack) > 0 && i < left && i < stack[0]; i++ {
		if arr[i] <= arr[stack[0]] {
			res = min(res, stack[0]-i-1)
		} else {
			stack = stack[1:]
		}
	}
	res = min(res, n-left-len(stack))
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
