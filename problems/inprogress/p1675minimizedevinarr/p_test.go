package p1675minimizedevinarr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumDeviation(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{9, 4, 3, 6, 2}, 7},
		// {[]int{1, 2, 3, 4}, 1},
		// {[]int{4, 1, 5, 20, 3}, 3},
		// {[]int{2, 10, 8}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minimumDeviation(tc.nums))
		})
	}
}

func Test_d(t *testing.T) {
	diff := d([]int{9, 4, 3, 6, 2})
	require.Equal(t, 7, diff)
}

func d(nums []int) int {
	var max int
	min := 1000000
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return max - min
}

func minimumDeviation(nums []int) int {
	return min2(nums, []int{})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min2(avail []int, used []int) int {
	if len(avail) == 0 {
		return d(used)
	}
	n := avail[0]
	minD := min2(avail[1:], append(used, n))
	if n%2 == 1 {
		minD = min(minD, min2(avail[1:], append(used, n*2)))
	} else {
		for n%2 == 0 {
			n /= 2
			minD = min(minD, min2(avail[1:], append(used, n)))
		}
	}
	return minD
}
