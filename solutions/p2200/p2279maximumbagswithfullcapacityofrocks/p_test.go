package p2279maximumbagswithfullcapacityofrocks

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumBags(t *testing.T) {
	for _, tc := range []struct {
		capacity        []int
		rocks           []int
		additionalRocks int
		want            int
	}{
		{
			[]int{91, 54, 63, 99, 24, 45, 78},
			[]int{35, 32, 45, 98, 6, 1, 25},
			17,
			1,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.capacity), func(t *testing.T) {
			require.Equal(t, tc.want, maximumBags(tc.capacity, tc.rocks, tc.additionalRocks))
		})
	}
}

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	h := IntHeap{}

	for i, c := range capacity {
		r := rocks[i]
		h = append(h, c-r)
	}
	heap.Init(&h)
	var count int
	for len(h) > 0 {
		x := heap.Pop(&h).(int)
		if x > additionalRocks {
			break
		}
		additionalRocks -= x
		count++
	}
	return count
}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
