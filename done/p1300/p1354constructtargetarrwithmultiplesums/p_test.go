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
	n := len(target)
	if n == 1 {
		return target[0] == 1
	}
	pq := make(PQ, n)
	var total int
	for i, n := range target {
		total += n
		pq[i] = &Item{idx: i, val: n}
	}
	heap.Init(&pq)
	for pq[0].val > 1 {
		rest := total - pq[0].val
		if rest == 1 {
			return true
		}
		x := pq[0].val % rest
		if x == 0 || x == pq[0].val {
			return false
		}
		total = total - pq[0].val + x
		pq[0].val = x
		heap.Fix(&pq, pq[0].idx)
	}

	for _, item := range pq {
		if item.val != 1 {
			return false
		}
	}
	return true
}

type Item struct {
	idx int
	val int
}

type PQ []*Item

func (h PQ) Len() int { return len(h) }
func (h PQ) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}
func (h PQ) Less(i, j int) bool {
	return h[i].val > h[j].val
}
func (h *PQ) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}
func (h *PQ) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
