package p0845longestmountaininarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestMountain(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{0, 2, 0, 2, 1, 2, 3, 4, 4, 1}, 3},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 0},
		{[]int{3, 2}, 0},
		{[]int{2, 1, 4, 7, 3, 2, 5}, 5},
		{[]int{2, 2, 2}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, longestMountain(tc.arr))
		})
	}
}

func longestMountain(arr []int) int {
	n := len(arr)
	var maxLen int
	// Let's detect peaks then just scan on both sides
	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1] >= arr[i] || arr[i] <= arr[i+1] {
			continue
		}
		// there's a peak here
		l := i - 1
		for l-1 >= 0 && arr[l-1] < arr[l] {
			l--
		}
		r := i + 1
		for r+1 < n && arr[r+1] < arr[r] {
			r++
		}
		maxLen = max(maxLen, r-l+1)
		i = r
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
