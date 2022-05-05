package p0978longestturbulentsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxTurbulenceSize(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24}, 8},
		{[]int{9, 4, 2, 10, 7, 8, 8, 1, 9}, 5},
		{[]int{4, 8, 12, 16}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, maxTurbulenceSize(tc.arr))
		})
	}
}

func maxTurbulenceSize(arr []int) int {
	n := len(arr)

	maxLen := 1
	for _, pos := range [][2]int{{0, 1}, {1, 0}} {
		gtPos := pos[0]
		ltPos := pos[1]
		start := 0
		for i := 1; i < n; i++ {
			maxLen = max(maxLen, i-start)
			if arr[i] == arr[i-1] {
				start = i
				continue
			}
			if i%2 == ltPos && arr[i] > arr[i-1] ||
				i%2 == gtPos && arr[i] < arr[i-1] {
				// Not OK, change start to next position
				start = i
				continue
			}
		}
		maxLen = max(maxLen, n-start)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
