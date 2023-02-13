package p2558takegiftsfromtherichestpile

import (
	"container/heap"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pickGifts(t *testing.T) {
	for i, tc := range []struct {
		gifts []int
		k     int
		want  int64
	}{
		{[]int{25, 64, 9, 4, 100}, 4, 29},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, pickGifts(tc.gifts, tc.k))
		})
	}
}

func pickGifts(gifts []int, k int) int64 {
	var res int64
	var h maxHeap
	for i := range gifts {
		h = append(h, gifts[i])
	}
	heap.Init(&h)
	for i := 0; i < k; i++ {
		a := sqrt(h[0])
		h[0] = int(a)
		heap.Fix(&h, 0)
	}
	for i := range h {
		res += int64(h[i])
	}
	return res
}

func sqrt(x int) int {
	lo, hi := 1, math.MaxInt32
	for lo < hi {
		mid := lo + (hi-lo)/2
		if mid*mid <= x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}

type maxHeap []int

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].heapIdx = i
	// h[j].heapIdx = j
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *maxHeap) Push(x interface{}) {
	el := x.(int)
	// el.heapIdx = len(*h)
	*h = append(*h, el)
}
func (h *maxHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	// el = nil
	*h = (*h)[:n-1]
	return el
}
