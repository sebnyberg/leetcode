package p3346maximumfrequencyofanelementafterperformingoperationsi

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxFrequency(t *testing.T) {
	for _, tc := range []struct {
		nums          []int
		k             int
		numOperations int
		want          int
	}{
		{[]int{1, 2, 4, 5}, 2, 4, 4},
		{[]int{1, 4, 5}, 1, 2, 2},
		{[]int{9}, 0, 0, 1},
		{[]int{5, 11, 20, 20}, 5, 1, 2},
		{[]int{1, 4, 5}, 1, 2, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxFrequency(tc.nums, tc.k, tc.numOperations))
		})
	}
}

func maxFrequency(nums []int, k int, numOperations int) int {
	type delta struct {
		i int
		d int
	}
	var deltas []delta
	for _, x := range nums {
		deltas = append(deltas,
			delta{x - k, 1},
			delta{x, 0},
			delta{x + k + 1, -1},
		)
	}
	sort.Slice(deltas, func(i, j int) bool {
		return deltas[i].i < deltas[j].i
	})
	var intervals int
	var j int
	var res int
	for j < len(deltas) {
		var middleCount int
		if deltas[j].d == 0 {
			middleCount++
		} else {
			intervals += deltas[j].d
		}
		j++
		for j < len(deltas) && deltas[j].i == deltas[j-1].i {
			if deltas[j].d == 0 {
				middleCount++
			} else {
				intervals += deltas[j].d
			}
			j++
		}
		res = max(res, middleCount+max(min(intervals-middleCount, numOperations), 0))
	}
	return res
}
