package p2158amountofnewareapaintedeachday

import (
	"container/heap"
	"fmt"
	"github.com/sebnyberg/leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_amountPainted(t *testing.T) {
	for _, tc := range []struct {
		paint [][]int
		want  []int
	}{
		{leetcode.ParseMatrix("[[1,4],[4,7],[5,8]]"), []int{3, 3, 1}},
		{leetcode.ParseMatrix("[[1,4],[5,8],[4,7]]"), []int{3, 3, 1}},
		{leetcode.ParseMatrix("[[1,5],[2,4]]"), []int{4, 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.paint), func(t *testing.T) {
			require.Equal(t, tc.want, amountPainted(tc.paint))
		})
	}
}

func amountPainted(paint [][]int) []int {
	type paintStrip struct {
		day     int
		isStart bool
	}
	n := len(paint)
	var maxPos int
	for _, p := range paint {
		if p[1] > maxPos {
			maxPos = p[1]
		}
	}
	painting := make([][]paintStrip, maxPos+1)
	for i, p := range paint {
		painting[p[0]] = append(painting[p[0]], paintStrip{i, true})
		painting[p[1]] = append(painting[p[1]], paintStrip{i, false})
	}
	layers := &LayerHeap{
		layers:       make([]int, 0, n),
		layerHeapIdx: make([]int, n+1),
	}
	res := make([]int, n)
	// For each position in the painting
	for i := 0; i < maxPos; i++ {

		// Visit all strips starting / ending in this position
		for _, strip := range painting[i] {
			if strip.isStart {
				layers.Add(strip.day)
			} else {
				layers.Remove(strip.day)
			}
		}

		// Attribute this paint position to the smallest day of painting
		if smallestDay := layers.GetMin(); smallestDay != -1 {
			res[smallestDay]++
		}
	}
	return res
}

type LayerHeap struct {
	layers       []int
	layerHeapIdx []int
}

func (h *LayerHeap) GetMin() int {
	if len(h.layers) == 0 {
		return -1
	}
	return h.layers[0]
}
func (h *LayerHeap) Add(x int) {
	heap.Push(h, x)
}
func (h *LayerHeap) Remove(x int) {
	idx := h.layerHeapIdx[x]
	heap.Remove(h, idx)
}

func (h *LayerHeap) Len() int { return len(h.layers) }
func (h *LayerHeap) Swap(i, j int) {
	h.layerHeapIdx[h.layers[i]] = j
	h.layerHeapIdx[h.layers[j]] = i
	h.layers[i], h.layers[j] = h.layers[j], h.layers[i]
}
func (h *LayerHeap) Less(i, j int) bool {
	return h.layers[i] < h.layers[j]
}
func (h *LayerHeap) Push(x interface{}) {
	h.layerHeapIdx[x.(int)] = len(h.layers)
	h.layers = append(h.layers, x.(int))
}
func (h *LayerHeap) Pop() interface{} {
	n := len(h.layers)
	item := h.layers[n-1]
	h.layers = h.layers[:n-1]
	return item
}
