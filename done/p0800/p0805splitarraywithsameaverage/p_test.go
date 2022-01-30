package p0805splitarraywithsameaverage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_splitArraySameAverage(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, true},
		{[]int{3, 1}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, splitArraySameAverage(tc.nums))
		})
	}
}

func splitArraySameAverage(nums []int) bool {
	n := len(nums)
	var sum int
	for _, n := range nums {
		sum += n
	}

	sums := make([]int, sum/2+1)
	for _, num := range nums {
		if num > sum/2 {
			continue
		}
		for s := sum / 2; s > num; s-- {
			if sums[s-num] != 0 {
				sums[s] |= sums[s-num] << 1
			}
		}
		sums[num] = (sums[num] | sums[0]<<1) | 1
	}

	for k := 1; k <= n/2; k++ {
		if (sum*k)%n == 0 && sums[(sum*k)/n]&(1<<(k-1)) != 0 {
			return true
		}
	}
	return false
}
