package p0308rangesumquery2dmutable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumArray(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		nm := Constructor([]int{1, 3, 5})
		res := nm.SumRange(0, 2)
		require.Equal(t, 9, res)
		nm.Update(1, 2)
		res = nm.SumRange(0, 2)
		require.Equal(t, 8, res)
	})

}

type NumArray struct {
	n    int
	bit  []int32
	nums []int32
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	nm := NumArray{n, make([]int32, n+1), make([]int32, n)}
	for i, n := range nums {
		nm.nums[i] = int32(n)
	}
	copy(nm.bit[1:], nm.nums)
	for i := 1; i < n; i++ {
		parent := i + (i & -i)
		if parent <= n {
			nm.bit[parent] += nm.bit[i]
		}
	}
	return nm
}

func (this *NumArray) Update(index int, val int) {
	i := index + 1
	d := this.nums[index] - int32(val)
	this.nums[index] = int32(val)
	for i < len(this.bit) {
		this.bit[i] -= d
		i += i & -i
	}
}

func (this *NumArray) Sum(i int) int {
	var res int
	for i > 0 {
		res += int(this.bit[i])
		i -= i & -i
	}
	return res
}

func (this *NumArray) SumRange(left int, right int) int {
	leftSum := this.Sum(left)
	rightSum := this.Sum(right + 1)
	return rightSum - leftSum
}
