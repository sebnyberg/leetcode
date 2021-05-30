package p1167mincosttoconnectsticks

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_connectSticks(t *testing.T) {
	for _, tc := range []struct {
		sticks []int
		want   int
	}{
		{[]int{2, 4, 3}, 14},
		{[]int{1, 8, 3, 5}, 30},
		{[]int{5}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.sticks), func(t *testing.T) {
			require.Equal(t, tc.want, connectSticks(tc.sticks))
		})
	}
}

func connectSticks(sticks []int) int {
	h := IntHeap(sticks)
	heap.Init(&h)
	var res int
	for len(h) > 1 {
		newStick := heap.Pop(&h).(int) + heap.Pop(&h).(int)
		res += newStick
		heap.Push(&h, newStick)
	}
	return res
}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
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
