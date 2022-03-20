package p2208

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_halveArray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{5, 19, 8, 1}, 3},
		{[]int{3, 8, 20}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, halveArray(tc.nums))
		})
	}
}

func halveArray(nums []int) int {
	n := len(nums)
	h := make(FloatHeap, n)
	var sum float64
	for i, v := range nums {
		h[i] = float64(v)
		sum += float64(v)
	}
	heap.Init(&h)
	curr := sum
	half := sum / 2
	var ops int
	for curr > half {
		x := heap.Pop(&h).(float64)
		xx := x / 2
		curr -= xx
		heap.Push(&h, xx)
		ops++
	}
	return ops
}

type FloatHeap []float64

func (h FloatHeap) Len() int { return len(h) }
func (h FloatHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h FloatHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *FloatHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}
func (h *FloatHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
