package p2402meetingroomsiii

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_mostBooked(t *testing.T) {
	for _, tc := range []struct {
		n        int
		meetings [][]int
		want     int
	}{
		{2, leetcode.ParseMatrix("[[4,11],[1,13],[8,15],[9,18],[0,17]]"), 1},
		{4, leetcode.ParseMatrix("[[18,19],[3,12],[17,19],[2,13],[7,10]]"), 0},
		{2, leetcode.ParseMatrix("[[0,10],[1,5],[2,7],[3,4]]"), 0},
		{3, leetcode.ParseMatrix("[[1,20],[2,10],[3,5],[4,9],[6,8]]"), 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, mostBooked(tc.n, tc.meetings))
		})
	}
}

func mostBooked(n int, meetings [][]int) int {
	// Use a priority queue to keep track of the most eligible room for a given
	// meeting.
	h := make(roomHeap, n)
	for i := range h {
		h[i] = room{
			idx:         i,
			booked:      0,
			availableAt: 0,
		}
	}
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	var maxBooked, maxBookedIdx int
	for _, m := range meetings {
		for h[0].availableAt < m[0] {
			h[0].availableAt = m[0]
			heap.Fix(&h, 0)
		}
		h[0].booked++
		h[0].availableAt += m[1] - m[0]
		if h[0].booked > maxBooked || h[0].booked == maxBooked && h[0].idx < maxBookedIdx {
			maxBooked = h[0].booked
			maxBookedIdx = h[0].idx
		}
		heap.Fix(&h, 0)
	}
	return maxBookedIdx
}

type room struct {
	idx         int
	booked      int
	availableAt int
}

type roomHeap []room

func (h roomHeap) Len() int      { return len(h) }
func (h roomHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h roomHeap) Less(i, j int) bool {
	if h[i].availableAt == h[j].availableAt {
		return h[i].idx < h[j].idx
	}
	return h[i].availableAt < h[j].availableAt
}
func (h *roomHeap) Push(x interface{}) { *h = append(*h, x.(room)) }
func (h *roomHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
