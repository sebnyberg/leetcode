package p1852distinctnumbersineachsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_distinctNumbers(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 2, 2, 1, 3}, 3, []int{3, 2, 2, 2, 3}},
		{[]int{1, 1, 1, 1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, distinctNumbers(tc.nums, tc.k))
		})
	}
}

func distinctNumbers(nums []int, k int) []int {
	var count [100001]int
	var ndistinct int
	for i := 0; i < k; i++ {
		count[nums[i]]++
		if count[nums[i]] == 1 {
			ndistinct++
		}
	}

	n := len(nums)
	res := make([]int, n-k+1)
	res[0] = ndistinct
	for i := k; i < n; i++ {
		count[nums[i]]++
		if count[nums[i]] == 1 {
			ndistinct++
		}
		count[nums[i-k]]--
		if count[nums[i-k]] == 0 {
			ndistinct--
		}
		res[i-k+1] = ndistinct
	}
	return res
}
