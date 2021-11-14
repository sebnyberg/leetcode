package p2070mostbeautifulitemforeachquery

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumBeauty(t *testing.T) {
	for _, tc := range []struct {
		items   [][]int
		queries []int
		want    []int
	}{
		{[][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}, []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 5, 5, 6, 6}},
		{[][]int{{1, 2}, {1, 2}, {1, 3}, {1, 4}}, []int{1}, []int{4}},
		{[][]int{{10, 1000}}, []int{5}, []int{0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.items), func(t *testing.T) {
			require.Equal(t, tc.want, maximumBeauty(tc.items, tc.queries))
		})
	}
}

func maximumBeauty(items [][]int, queries []int) []int {
	// Sort by price
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	// Iterate over items, adding only items which increase the beauty
	var trimmedLen int
	var maxBeauty int
	for j, item := range items {
		if item[1] < maxBeauty { // skip
			continue
		}
		maxBeauty = item[1]
		items[trimmedLen] = items[j]
		trimmedLen++
	}
	items = items[:trimmedLen]

	res := make([]int, len(queries))
	for i, q := range queries {
		abovePriceIdx := sort.Search(len(items), func(j int) bool {
			return items[j][0] > q
		})
		if abovePriceIdx == 0 {
			continue
		}
		res[i] = items[abovePriceIdx-1][1]
	}

	return res
}
