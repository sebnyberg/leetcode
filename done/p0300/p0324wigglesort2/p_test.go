package p0324wigglesort2

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wiggleSort(t *testing.T) {
	for _, tc := range []struct {
		nums []int
	}{
		{[]int{4, 5, 5, 6}},
		{[]int{1, 5, 1, 1, 6, 4}},
		{[]int{1, 3, 2, 2, 3, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			wiggleSort(tc.nums)
			for i := range tc.nums {
				if i == 0 {
					continue
				}
				if i%2 == 0 {
					require.True(t, tc.nums[i] < tc.nums[i-1])
				} else {
					require.True(t, tc.nums[i] > tc.nums[i-1])
				}
			}
		})
	}
}

func wiggleSort(nums []int) {
	n := len(nums)
	sortedNums := make([]int, n)
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	half := (n + 1) / 2
	for i := range sortedNums {
		if i%2 == 0 {
			nums[i] = sortedNums[half-1-i/2]
		} else {
			nums[i] = sortedNums[n-1-i/2]
		}
	}
}
