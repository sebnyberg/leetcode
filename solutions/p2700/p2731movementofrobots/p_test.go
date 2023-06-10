package p2731movementofrobots

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumDistance(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		s    string
		d    int
		want int
	}{
		{[]int{-2, 0, 2}, "RLL", 3, 8},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, sumDistance(tc.nums, tc.s, tc.d))
		})
	}
}

const mod = 1e9 + 7

func sumDistance(nums []int, s string, d int) int {
	// It seems like each number is just passing through the other.
	//
	// So we just need to calculate distances from numbers after adding /
	// removing D from each value
	n := len(nums)
	res := make([]int, n)
	for i := range nums {
		if s[i] == 'L' {
			res[i] = nums[i] - d
		} else {
			res[i] = nums[i] + d
		}
	}
	sort.Ints(res)
	var sum int
	var final int
	for i := 1; i < len(res); i++ {
		ddd := res[i] - res[i-1]
		sum = (sum + ddd*i) % mod
		final = (final + sum) % mod
	}
	return final
}
