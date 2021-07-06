package p0327countofrangesum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countRangeSum(t *testing.T) {
	for _, tc := range []struct {
		nums         []int
		lower, upper int
		want         int
	}{
		{[]int{5, -23, -5, -1, -21, 13, 15, 7, 18, 4, 7, 26, 29, -7, -28, 11, -20, -29, 19, 22, 15, 25, 17, -13, 11, -15, 19, -8, 3, 12, -1, 2, -1, -21, -10, -7, 14, -12, -14, -8, -1, -30, 19, -27, 16, 2, -15, 23, 6, 14, 23, 2, -4, 4, -9, -8, 10, 20, -29, 29},
			-19,
			10,
			362,
		},
		{[]int{-2, 5, -1}, -2, 2, 3},
		{[]int{0}, 0, 0, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countRangeSum(tc.nums, tc.lower, tc.upper))
		})
	}
}

func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
	presum := make([]int, n+1)
	for i, num := range nums {
		presum[i+1] = presum[i] + num
	}
	return merge(presum, 0, n, lower, upper)
}

func merge(sum []int, l, r, lower, upper int) int {
	if l >= r {
		return 0
	}
	mid := (l + r) / 2
	count := 0
	count += merge(sum, l, mid, lower, upper) + merge(sum, mid+1, r, lower, upper)
	m, n := mid+1, mid+1
	for i := l; i <= mid; i++ {
		for m <= r && sum[m]-sum[i] < lower {
			m++
		}
		for n <= r && sum[n]-sum[i] <= upper {
			n++
		}
		count += n - m
	}
	left := append([]int{}, sum[l:mid+1]...)
	right := append([]int{}, sum[mid+1:r+1]...)
	i := l
	ll := 0
	rr := 0
	for ll < len(left) && rr < len(right) {
		if left[ll] < right[rr] {
			sum[i] = left[ll]
			ll++
		} else {
			sum[i] = right[rr]
			rr++
		}
		i++
	}
	for rr < len(right) {
		sum[i] = right[rr]
		rr++
		i++
	}
	for ll < len(left) {
		sum[i] = left[ll]
		ll++
		i++
	}
	return count
}
