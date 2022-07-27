package p2341maxsumofapariwithequalsumofdigits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{18, 43, 36, 13, 7}, 54},
		{[]int{10, 12, 19, 13}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maximumSum(tc.nums))
		})
	}
}

func maximumSum(nums []int) int {
	maxSumVal := make(map[int]int)
	digitSum := func(x int) int {
		var res int
		for x > 0 {
			res += x % 10
			x /= 10
		}
		return res
	}
	var res int
	for _, x := range nums {
		s := digitSum(x)
		v, exists := maxSumVal[s]
		if exists && v+x > res {
			res = v + x
		}
		maxSumVal[s] = max(v, x)
	}
	if res == 0 {
		return -1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
