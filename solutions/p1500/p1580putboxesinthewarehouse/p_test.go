package p1580putboxesinthewarehouse

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
		{[]int{1, 2, 2, 3, 4}, []int{3, 4, 1, 2}, 4},
		{[]int{4, 5, 6, 2}, []int{3, 4, 1, 2}, 1},
		{[]int{3, 5, 5, 2}, []int{2, 1, 3, 4, 5}, 3},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, 3},
		{[]int{4, 5, 6}, []int{3, 3, 3, 3, 3}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.boxes), func(t *testing.T) {
			require.Equal(t, tc.want, maxBoxesInWarehouse(tc.boxes, tc.warehouse))
		})
	}
}

func maxBoxesInWarehouse(boxes []int, warehouse []int) int {
	// Find min point in the warehouse
	minVal := warehouse[0]
	var minIdx int
	for i, n := range warehouse {
		if n < minVal {
			minVal = n
			minIdx = i
		}
	}

	// From the right edge to the min Idx, take the min value of the current
	// and the previous
	nw := len(warehouse)
	for i := nw - 2; i >= minIdx; i-- {
		warehouse[i] = min(warehouse[i], warehouse[i+1])
	}

	// From the left edge to min Idx, take the min value of cur and prev
	for i := 1; i <= minIdx; i++ {
		warehouse[i] = min(warehouse[i], warehouse[i-1])
	}

	// Sort boxes descending order
	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i] > boxes[j]
	})

	// Keep three cursors, one for the minimum box, one for left warehouse side,
	// one for right warehouse side
	b, wl, wr := len(boxes)-1, minIdx-1, minIdx
	var nboxes int
	for b >= 0 && (wl >= 0 || wr <= nw-1) {
		// Pick warehouse side - always prefer lowest height
		switch {
		case wl >= 0 && (wr == nw || warehouse[wl] <= warehouse[wr]):
			if boxes[b] <= warehouse[wl] {
				b--
				wl--
				nboxes++
			} else {
				wl--
			}
		case wr <= nw-1 && (wl == -1 || warehouse[wr] < warehouse[wl]):
			if boxes[b] <= warehouse[wr] {
				b--
				wr++
				nboxes++
			} else {
				wr++
			}
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
