package p1887reductionoperationstomakethearrayelementsequal

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reductionOperations(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{5, 1, 3}, 3},
		{[]int{1, 1, 1}, 0},
		{[]int{1, 1, 2, 2, 3}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, reductionOperations(tc.nums))
		})
	}
}

func reductionOperations(nums []int) int {
	var numFreq [50001]int
	for _, num := range nums {
		numFreq[num]++
	}
	h := make(FreqHeap, 0)
	for num, count := range numFreq {
		if count > 0 {
			h = append(h, &Freq{num, count})
		}
	}
	heap.Init(&h)
	var res int
	for len(h) > 1 {
		// Pop highest value off stack
		el := heap.Pop(&h).(*Freq)
		res += el.count
		h[0].count += el.count
	}
	return res
}

type Freq struct {
	val   int
	count int
}

type FreqHeap []*Freq

func (h FreqHeap) Len() int { return len(h) }
func (h FreqHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h FreqHeap) Less(i, j int) bool {
	return h[i].val > h[j].val
}
func (h *FreqHeap) Push(x interface{}) {
	*h = append(*h, x.(*Freq))
}
func (h *FreqHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
