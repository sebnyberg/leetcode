package p1649createsortedarrinst

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_createSortedArray(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want int
	}{
		{[]int{1, 5, 6, 2}, 1},
		{[]int{1, 2, 3, 6, 5, 4}, 3},
		{[]int{1, 3, 3, 3, 2, 4, 2, 1, 2}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, createSortedArray(tc.in))
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const mod = 1e9 + 7

func createSortedArray(instructions []int) (cost int) {
	nums := make([]int, len(instructions))

	for i, n := range instructions {
		insertPos := sort.SearchInts(nums[:i], n)
		nextPos := sort.SearchInts(nums[insertPos:i], n+1)

		cost += min(insertPos, i-(insertPos+nextPos))

		copy(nums[insertPos+1:i+1], nums[insertPos:i])
		nums[insertPos] = n
	}

	return cost % mod
}
