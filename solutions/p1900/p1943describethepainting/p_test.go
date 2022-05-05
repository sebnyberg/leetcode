package p1943describethepainting

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_splitPainting(t *testing.T) {
	for _, tc := range []struct {
		segments [][]int
		want     [][]int64
	}{
		{[][]int{{1, 7, 9}, {6, 8, 15}, {8, 10, 7}}, [][]int64{{1, 6, 9}, {6, 7, 24}, {7, 8, 15}, {8, 10, 7}}},
		{[][]int{{1, 4, 5}, {1, 4, 7}, {4, 7, 1}, {4, 7, 11}}, [][]int64{{1, 4, 12}, {4, 7, 12}}},
		{[][]int{{4, 16, 12}, {9, 10, 15}, {18, 19, 13}, {3, 13, 20}, {12, 16, 3}, {2, 10, 10}, {3, 11, 4}, {13, 16, 6}}, [][]int64{{2, 3, 10}, {3, 4, 34}, {4, 9, 46}, {9, 10, 61}, {10, 11, 36}, {11, 12, 32}, {12, 13, 35}, {13, 16, 21}, {18, 19, 13}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.segments), func(t *testing.T) {
			require.Equal(t, tc.want, splitPainting(tc.segments))
		})
	}
}

func splitPainting(segments [][]int) [][]int64 {
	h := make(ColorHeap, 0)
	for _, seg := range segments {
		start, end, color := seg[0], seg[1], seg[2]
		h = append(h, coloring{start, color})
		h = append(h, coloring{end, -color})
	}
	heap.Init(&h)
	res := []coloredSegment{{}}
	var resIdx int
	var currentVal int
	for len(h) > 0 {
		// Collect all segments which start (or ended before) this position
		x := heap.Pop(&h).(coloring)
		currentVal += x.val
		for len(h) > 0 && h[0].pos == x.pos {
			currentVal += h[0].val
			heap.Pop(&h)
		}
		res[resIdx].end = x.pos
		resIdx++
		res = append(res, coloredSegment{x.pos, -1, currentVal})
	}
	ret := make([][]int64, 0, len(res)-2)
	for _, el := range res[1 : len(res)-1] {
		if el.val > 0 {
			ret = append(ret, []int64{int64(el.start), int64(el.end), int64(el.val)})
		}
	}
	return ret
}

type coloredSegment struct {
	start, end, val int
}

type coloring struct {
	pos, val int
}

type ColorHeap []coloring

func (h ColorHeap) Len() int { return len(h) }
func (h ColorHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h ColorHeap) Less(i, j int) bool {
	return h[i].pos < h[j].pos
}
func (h *ColorHeap) Push(x interface{}) {
	*h = append(*h, x.(coloring))
}
func (h *ColorHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
