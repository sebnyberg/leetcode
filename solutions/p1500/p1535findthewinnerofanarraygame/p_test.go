package p1535findthewinnerofanarraygame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getWinner(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		k    int
		want int
	}{
		{[]int{2, 1, 3, 5, 4, 6, 7}, 2, 5},
		{[]int{1, 25, 35, 42, 68, 70}, 1, 25},
		{[]int{3, 2, 1}, 10, 3},
		{[]int{1, 9, 8, 2, 3, 7, 6, 4, 5}, 7, 9},
		{[]int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99}, 1000000000, 99},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, getWinner(tc.arr, tc.k))
		})
	}
}

func getWinner(arr []int, k int) int {
	// Since the elements in arr are distinct, there will always be a winner,
	// no matter how big k is.
	// The re-shuffling is basically a sorting of arr.
	var cur int
	if arr[0] > arr[1] {
		cur = arr[0]
	} else {
		cur = arr[1]
	}
	count := 1
	maxNum := cur
	for _, n := range arr[2:] {
		if count == k {
			return cur
		}
		maxNum = max(maxNum, n)
		if n < cur {
			count++
		} else {
			count = 1
			cur = n
		}
	}
	return maxNum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
