package p0280wigglesort

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wiggleSort(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		// {[]int{3, 5, 2, 1, 6, 4}, []int{3, 5, 1, 6, 2, 4}},
		// {[]int{6, 6, 5, 6, 3, 8}, []int{6, 6, 5, 6, 3, 8}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			wiggleSort(tc.nums)
			require.Equal(t, tc.want, tc.nums)
		})
	}
}

func wiggleSort(nums []int) {
	n := len(nums)
	sortedNums := make([]int, n)
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	for i := range sortedNums {
		if i%2 == 0 {
			nums[i] = sortedNums[i/2]
		} else {
			nums[i] = sortedNums[n-1-(i/2)]
		}
	}
}
