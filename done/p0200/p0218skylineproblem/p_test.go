package p0218skylineproblem

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getSkyline(t *testing.T) {
	for _, tc := range []struct {
		buildings [][]int
		want      [][]int
	}{
		{
			[][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}},
			[][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.buildings), func(t *testing.T) {
			require.Equal(t, tc.want, getSkyline(tc.buildings))
		})
	}
}

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	edges := make([]*Building, n*2)
	for i, b := range buildings {
		edges[2*i] = &Building{b[0], b[2], i, 0}
		edges[2*i+1] = &Building{b[1], -b[2], i, 0}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].x < edges[j].x
	})
	h := &MaxHeap{}
	heap.Init(h)
	heap.Push(h, &Building{})
	inserts := make(map[int]*Building)
	ret := make([][]int, 0)
	for i, edge := range edges {
		if edge.height >= 0 { // left side, add to maxheap
			edge.heapPosition = h.Len()
			inserts[edge.originalIndex] = edge
			heap.Push(h, edge)
		} else {
			heap.Remove(h, inserts[edge.originalIndex].heapPosition)
			delete(inserts, edge.originalIndex)
		}
		if i == n*2-1 || edge.x != edges[i+1].x {
			currMax := (*h)[0].height
			if len(ret) == 0 || currMax != ret[len(ret)-1][1] {
				ret = append(ret, []int{edge.x, currMax})
			}
			if i == n*2-1 {
				break
			}
		}
	}

	return ret
}

type Building struct {
	x, height, originalIndex, heapPosition int
}

type MaxHeap []*Building // x, y, index in heap
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapPosition, h[j].heapPosition = i, j
}
func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h *MaxHeap) Push(v interface{}) { *h = append(*h, v.(*Building)) }
func (h *MaxHeap) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}
