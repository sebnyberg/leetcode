package p0078subsets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subsets(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, subsets(tc.nums))
		})
	}
}

func subsets(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 1<<n)
	for x := 0; x < 1<<n; x++ {
		res[x] = make([]int, 0, n)
		for b := 0; b < n; b++ {
			if x&(1<<b) > 0 {
				res[x] = append(res[x], nums[b])
			}
		}
	}
	return res
}
