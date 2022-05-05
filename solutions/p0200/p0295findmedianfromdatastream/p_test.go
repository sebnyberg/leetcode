package p0295findmedianfromdatastream

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMedianFinder(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		mf := Constructor()
		mf.AddNum(6)
		res := mf.FindMedian()
		require.InDelta(t, 6.0, res, 0.0001)
		mf.AddNum(10)
		res = mf.FindMedian()
		require.InDelta(t, 8.0, res, 0.0001)
		mf.AddNum(2)
		res = mf.FindMedian()
		require.InDelta(t, 6.0, res, 0.0001)
	})
}

type MedianFinder struct {
	nums []int
	n    int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		nums: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	idx := sort.SearchInts(this.nums, num)
	this.nums = append(this.nums, num)
	if idx != this.n {
		copy(this.nums[idx+1:], this.nums[idx:])
		this.nums[idx] = num
	}
	this.n++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.n == 0 {
		return 0
	}
	m := this.n / 2
	if this.n%2 == 0 {
		return float64(this.nums[m-1]+this.nums[m]) / 2
	}
	return float64(this.nums[m])
}
