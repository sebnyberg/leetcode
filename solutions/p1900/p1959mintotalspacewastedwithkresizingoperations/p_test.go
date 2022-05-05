package p1959mintotalspacewastedwithkresizingoperations

import (
	"fmt"
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
	// Todo: implement
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
