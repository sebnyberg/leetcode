package p1331ranktransformofanarray

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_arrayRankTransform(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want []int
	}{
		{[]int{40, 10, 20, 30}, []int{4, 1, 2, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, arrayRankTransform(tc.arr))
		})
	}
}

func arrayRankTransform(arr []int) []int {
	n := len(arr)
	res := make([]int, n)
	items := make([]arrItem, n)
	for i, n := range arr {
		items[i] = arrItem{i, n}
	}
	sort.Slice(items, func(i, j int) bool { return items[i].val < items[j].val })
	rank := 1
	for i := 0; i < n; {
		var j int
		for j = i + 1; j < n && items[i].val == items[j].val; j++ {
		}
		for _, it := range items[i:j] {
			res[it.idx] = rank
		}
		rank++
		i = j
	}
	return res
}

type arrItem struct{ idx, val int }
