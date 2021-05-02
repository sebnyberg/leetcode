package p0303rangesumquery

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumArr(t *testing.T) {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})
	res := na.SumRange(0, 2)
	require.Equal(t, 1, res)
	res = na.SumRange(2, 5)
	require.Equal(t, -1, res)
	res = na.SumRange(0, 5)
	require.Equal(t, -3, res)
}

type NumArray struct {
	prefixSum []int
	nums      []int
	n         int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	na := NumArray{
		prefixSum: make([]int, n),
		nums:      nums,
		n:         n,
	}
	for i := range nums {
		if i == 0 {
			na.prefixSum[i] = nums[i]
			continue
		}
		na.prefixSum[i] = na.prefixSum[i-1] + nums[i]
	}
	return na
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right] - this.prefixSum[left] + this.nums[left]
}
