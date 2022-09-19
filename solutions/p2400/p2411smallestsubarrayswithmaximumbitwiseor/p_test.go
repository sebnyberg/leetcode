package p2411smallestsubarrayswithmaximumbitwiseor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{8, 10, 8}, []int{2, 1, 1}},
		{[]int{1, 0, 2, 1, 3}, []int{3, 3, 2, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, smallestSubarrays(tc.nums))
		})
	}
}

func smallestSubarrays(nums []int) []int {
	var bitCount [32]int
	canRemove := func(x int) bool {
		for b := 0; x > 0; b++ {
			if bitCount[b] == 1 && x&1 == 1 {
				return false
			}
			x >>= 1
		}
		return true
	}
	addNum := func(x int) {
		for b := 0; x > 0; b++ {
			bitCount[b] += x & 1
			x >>= 1
		}
	}
	removeNum := func(x int) {
		for b := 0; x > 0; b++ {
			bitCount[b] -= x & 1
			x >>= 1
		}
	}
	n := len(nums)
	r := n - 1
	res := make([]int, n)
	res[n-1] = 1
	addNum(nums[n-1])
	for l := n - 2; l >= 0; l-- {
		addNum(nums[l])
		// Remove as many numbers from right side as possible
		for r > l && canRemove(nums[r]) {
			removeNum(nums[r])
			r--
		}
		res[l] = r - l + 1
	}
	return res
}
