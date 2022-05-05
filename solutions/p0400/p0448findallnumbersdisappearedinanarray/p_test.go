package p0448findallnumbersdisappearedinanarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDisappearedNumbers(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
		{[]int{1, 1}, []int{2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findDisappearedNumbers(tc.nums))
		})
	}
}

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	seen := make([]bool, n)
	for _, n := range nums {
		seen[n-1] = true
	}
	res := make([]int, 0)
	for i, ok := range seen {
		if !ok {
			res = append(res, i+1)
		}
	}
	return res
}
