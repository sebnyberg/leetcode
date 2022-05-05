package p0090subset2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subsetsWithDup(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, subsetsWithDup(tc.nums))
		})
	}
}

func subsetsWithDup(nums []int) [][]int {
	available := make(map[int]int)
	for _, n := range nums {
		available[n]++
	}
	values := make([]int, 0)
	for k := range available {
		values = append(values, k)
	}

	res := make([][]int, 0)
	findSets(0, values, available, []int{}, &res)
	return res
}

func findSets(i int, values []int, available map[int]int, prefix []int, res *[][]int) {
	if i >= len(values) {
		*res = append(*res, prefix)
		return
	}
	for j := 0; j <= available[values[i]]; j++ {
		if j == available[values[i]] {
			// No copy needed for last iteration, simply append to prefix
			for k := 0; k < j; k++ {
				prefix = append(prefix, values[i])
			}
			findSets(i+1, values, available, prefix, res)
			break
		}
		prefixCpy := make([]int, len(prefix))
		copy(prefixCpy, prefix)
		for k := 0; k < j; k++ {
			prefixCpy = append(prefixCpy, values[i])
		}
		findSets(i+1, values, available, prefixCpy, res)
	}
}
