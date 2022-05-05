package p2190

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mostFrequent(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		key  int
		want int
	}{
		{[]int{1, 100, 200, 1, 100}, 1, 100},
		{[]int{2, 2, 2, 2, 3}, 2, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, mostFrequent(tc.nums, tc.key))
		})
	}
}

func mostFrequent(nums []int, key int) int {
	followCount := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			followCount[nums[i+1]]++
		}
	}
	var maxCount int
	var maxVal int
	for val, count := range followCount {
		if count > maxCount {
			maxVal = val
			maxCount = count
		}
	}
	return maxVal

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
