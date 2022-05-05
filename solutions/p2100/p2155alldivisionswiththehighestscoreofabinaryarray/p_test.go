package p2155alldivisionswiththehighestscoreofabinaryarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxScoreIndices(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{0, 0, 1, 0}, []int{2, 4}},
		{[]int{0, 0, 0}, []int{3}},
		{[]int{1, 1}, []int{0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxScoreIndices(tc.nums))
		})
	}
}

func maxScoreIndices(nums []int) []int {
	n := len(nums)
	ones := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if i < n-1 {
			ones[i] = ones[i+1]
		}
		if nums[i] == 1 {
			ones[i]++
		}
	}
	res := []int{0}
	maxVal := ones[0]
	var zeroes int
	for i := 1; i <= n; i++ {
		v := nums[i-1]
		if v == 0 {
			zeroes++
		}
		x := zeroes + ones[i]
		if x < maxVal {
			continue
		}
		if x > maxVal {
			res = res[:0]
			maxVal = x
		}
		res = append(res, i)
	}
	return res
}
