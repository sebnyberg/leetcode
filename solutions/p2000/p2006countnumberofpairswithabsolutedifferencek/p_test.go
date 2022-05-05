package p2006countnumberofpairswithabsolutedifferencek

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countKDifference(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 2, 1}, 1, 4},
		{[]int{1, 3}, 3, 0},
		{[]int{3, 2, 1, 5, 4}, 2, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countKDifference(tc.nums, tc.k))
		})
	}
}

func countKDifference(nums []int, k int) int {
	var count int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[j]-nums[i]) == k {
				count++
			}
		}
	}
	return count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
