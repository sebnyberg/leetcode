package p1354constructtargetarrwithmultiplesums

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPossible(t *testing.T) {
	for _, tc := range []struct {
		target []int
		want   bool
	}{
		{[]int{1, 1000000000}, true},
		{[]int{9, 9, 9}, false},
		{[]int{9, 3, 5}, true},
		{[]int{1, 1, 1, 2}, false},
		{[]int{8, 5}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, isPossible(tc.target))
		})
	}
}

func isPossible(target []int) bool {
	var sum int
	for _, t := range target {
		sum += t
	}

	h := maxHeap(target)
	heap.Init(&h)
	for {
		sum -= h[0]
		if h[0] == 1 || sum == 1 {
			return true
		}
		if sum >= h[0] || sum == 0 || h[0]%sum == 0 {
			return false
		}
		h[0] %= sum
		sum += h[0]
		heap.Fix(&h, 0)
	}
}

type maxHeap []int

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
