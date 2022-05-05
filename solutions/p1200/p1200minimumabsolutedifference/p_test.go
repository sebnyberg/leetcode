package p1200minimumabsolutedifference

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumAbsDifference(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want [][]int
	}{
		{[]int{-12, 17, -59, 50, 10, 83, 27, -79}, [][]int{{10, 17}}},
		{[]int{3, 8, -10, 23, 19, -4, -14, 27}, [][]int{{-14, -10}, {19, 23}, {23, 27}}},
		{[]int{1, 3, 6, 10, 15}, [][]int{{1, 3}}},
		{[]int{4, 2, 1, 3}, [][]int{{1, 2}, {2, 3}, {3, 4}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, minimumAbsDifference(tc.arr))
		})
	}
}

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minDist := math.MaxInt32
	res := make([][]int, 0)
	for i := 1; i < len(arr); i++ {
		a, b := arr[i-1], arr[i]
		d := abs(a - b)
		if d < minDist {
			res = res[:0]
			minDist = d
		}
		if d == minDist {
			res = append(res, []int{a, b})
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
