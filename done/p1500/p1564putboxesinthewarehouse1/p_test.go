package p1564putboxesinthewarehouse1

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxBoxesInWarehouse(t *testing.T) {
	for _, tc := range []struct {
		boxes     []int
		warehouse []int
		want      int
	}{
		{[]int{4, 3, 4, 1}, []int{5, 3, 3, 4, 1}, 3},
		{[]int{1, 2, 2, 3, 4}, []int{3, 4, 1, 2}, 3},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, 1},
		{[]int{4, 5, 6}, []int{3, 3, 3, 3, 3}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.boxes), func(t *testing.T) {
			require.Equal(t, tc.want, maxBoxesInWarehouse(tc.boxes, tc.warehouse))
		})
	}
}

func maxBoxesInWarehouse(boxes []int, warehouse []int) int {
	// Set warehouse heights to be capped by the lowest height from the left
	for i := range warehouse[1:] {
		warehouse[i+1] = min(warehouse[i+1], warehouse[i])
	}

	// Sort boxes descending order
	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i] > boxes[j]
	})

	// Keep two cursors, add boxes either until a cursor goes out of bounds
	b, w := len(boxes)-1, len(warehouse)-1
	var nboxes int
	for b >= 0 && w >= 0 {
		switch {
		case boxes[b] <= warehouse[w]:
			// Add the box
			b--
			w--
			nboxes++
		case boxes[b] > warehouse[w]:
			w--
		}
	}
	return nboxes
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
