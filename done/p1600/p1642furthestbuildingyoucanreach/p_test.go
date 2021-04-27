package p1642furthestbuildingyoucanreach

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_furthestBuilding(t *testing.T) {
	for _, tc := range []struct {
		heights []int
		bricks  int
		ladders int
		want    int
	}{
		{[]int{4, 2, 7, 6, 9, 14, 12}, 5, 1, 4},
		{[]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2, 7},
		{[]int{14, 3, 19, 3}, 17, 0, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.heights), func(t *testing.T) {
			require.Equal(t, tc.want, furthestBuilding(tc.heights, tc.bricks, tc.ladders))
		})
	}
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)
	var nbuildings int
	doInit := true
	h := make(MinHeap, 0, ladders)
	for i := 1; i < n; i++ {
		d := heights[i] - heights[i-1]
		if d <= 0 {
			nbuildings++
			continue
		}
		// Difference is non-zero
		// The maximum height differences should be managed with ladders
		if h.Len() < ladders {
			h = append(h, d)
			nbuildings++
			continue
		}
		if doInit && h.Len() == ladders {
			heap.Init(&h)
			doInit = false
		}
		// After this point, bricks are used until there are no more bricks
		heap.Push(&h, d)
		minDelta := heap.Pop(&h).(int)
		if minDelta > bricks {
			return nbuildings
		}
		bricks -= minDelta
		nbuildings++
	}
	return nbuildings
}

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
