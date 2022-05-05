package p1340jumpgamev

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxJumps(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		d    int
		want int
	}{
		{[]int{59, 8, 74, 27, 92, 36, 95, 78, 73, 54, 75, 37, 42, 15, 59, 84, 66, 25, 35, 61, 97, 16, 6, 52, 49, 18, 22, 70, 5, 59, 92, 85}, 20, 8},
		{[]int{3, 3, 3, 3, 3}, 3, 1},
		{[]int{7, 6, 5, 4, 3, 2, 1}, 1, 7},
		{[]int{7, 1, 7, 1, 7, 1}, 2, 2},
		{[]int{66}, 1, 1},
		{[]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, maxJumps(tc.arr, tc.d))
		})
	}
}

func maxJumps(arr []int, d int) int {
	type arrItem struct {
		idx int
		val int
	}
	n := len(arr)
	items := make([]arrItem, n)
	for i := range arr {
		items[i] = arrItem{i, arr[i]}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].val > items[j].val
	})
	maxJumps := make([]int, n)
	for _, it := range items {
		i := it.idx
		maxRight := arr[i]
		for r := i + 1; r < n && r-i <= d; r++ {
			if arr[r] > maxRight {
				maxRight = arr[r]
				maxJumps[i] = max(maxJumps[i], maxJumps[r]+1)
			}
		}
		maxLeft := arr[i]
		for l := i - 1; l >= 0 && i-l <= d; l-- {
			if arr[l] > maxLeft {
				maxLeft = arr[l]
				maxJumps[i] = max(maxJumps[i], maxJumps[l]+1)
			}
		}
	}

	var res int
	for _, jumps := range maxJumps {
		res = max(res, jumps)
	}
	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
