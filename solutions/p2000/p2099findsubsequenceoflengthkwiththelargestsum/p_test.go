package p2099findsubsequenceoflengthkwiththelargestsum

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSubsequence(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{2, 1, 3, 3}, 2, []int{3, 3}},
		{[]int{-1, -2, 3, 4}, 3, []int{-1, 3, 4}},
		{[]int{3, 4, 3, 3}, 2, []int{3, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxSubsequence(tc.nums, tc.k))
		})
	}
}

func maxSubsequence(nums []int, k int) []int {
	cpy := make([]int, len(nums))
	copy(cpy, nums)
	sort.Slice(cpy, func(i, j int) bool {
		return cpy[i] > cpy[j]
	})
	want := make(map[int]int)
	for i := 0; i < k; i++ {
		want[cpy[i]]++
	}
	res := make([]int, 0, k)
	for _, num := range nums {
		if want[num] > 0 {
			res = append(res, num)
			want[num]--
		}
	}
	return res
}
