package p1762buildingswithanoceanview

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findBuildings(t *testing.T) {
	for _, tc := range []struct {
		heights []int
		want    []int
	}{
		{[]int{4, 2, 3, 1}, []int{0, 2, 3}},
		{[]int{4, 3, 2, 1}, []int{0, 1, 2, 3}},
		{[]int{1, 3, 2, 4}, []int{3}},
		{[]int{2, 2, 2, 2}, []int{3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.heights), func(t *testing.T) {
			require.Equal(t, tc.want, findBuildings(tc.heights))
		})
	}
}

func findBuildings(heights []int) []int {
	n := len(heights)
	maxHeight := 0
	res := make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		if heights[i] > maxHeight {
			res = append(res, i)
			maxHeight = heights[i]
		}
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return res
}
