package p2113elementsinarrayafterremovingandreplacingelements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_elementInNums(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		queries [][]int
		want    []int
	}{
		{[]int{0, 1, 2}, [][]int{{0, 2}, {2, 0}, {3, 2}, {5, 0}}, []int{2, 2, -1, 0}},
		{[]int{2}, [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, []int{2, -1, 2, -1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, elementInNums(tc.nums, tc.queries))
		})
	}
}

func elementInNums(nums []int, queries [][]int) []int {
	n := len(nums)
	res := make([]int, len(queries))
	for i := range res {
		res[i] = -1 // default result
	}
	for i, q := range queries {
		t, idx := q[0], q[1]
		t %= n * 2
		if t < n && t+idx < n {
			res[i] = nums[t+idx]
		} else if t > n && idx < t-n {
			res[i] = nums[idx]
		}
	}
	return res
}
