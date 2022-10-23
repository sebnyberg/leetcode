package p2420findallgoodindices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_goodIndices(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{2, 1, 1, 1, 3, 4, 1}, 2, []int{2, 3}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, goodIndices(tc.nums, tc.k))
		})
	}
}

func goodIndices(nums []int, k int) []int {
	// Lets make this easy for us... just check
	n := len(nums)
	ok := make([]bool, n)
	var m int
	for i := n - 1; i > k; i-- {
		if i < n-1 && nums[i] <= nums[i+1] {
			m++
		} else {
			m = 1
		}
		if m >= k {
			ok[i-1] = true
		}
	}
	m = 0
	var res []int
	for i := 0; i < n-k; i++ {
		if i > 0 && nums[i-1] >= nums[i] {
			m++
		} else {
			m = 1
		}
		if m >= k && ok[i+1] {
			res = append(res, i+1)
		}
	}
	return res
}
