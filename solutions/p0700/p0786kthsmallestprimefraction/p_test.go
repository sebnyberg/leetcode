package p0786kthsmallestprimefraction

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthSmallestPrimeFraction(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 5}, 3, []int{2, 5}},
		{[]int{1, 7}, 1, []int{1, 7}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, kthSmallestPrimeFraction(tc.arr, tc.k))
		})
	}
}

type frac struct {
	i, j int
}

func (f frac) Less(other frac) bool {
	a := f.i * other.j
	b := other.i * f.j
	return a < b
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	h := make(itemHeap, len(arr)-1)
	for i := 0; i < n-1; i++ {
		h[i] = item{i, n - 1, frac{arr[i], arr[n-1]}}
	}
	heap.Init(&h)
	for kk := 1; kk < k; kk++ {
		if h[0].j-1 == h[0].i {
			// No more fractions
			heap.Pop(&h)
			continue
		}
		h[0].j--
		h[0].val.j = arr[h[0].j]
		heap.Fix(&h, 0)
	}
	res := []int{h[0].val.i, h[0].val.j}
	return res
}

type item struct {
	i, j int
	val  frac
}

type itemHeap []item

func (h itemHeap) Len() int { return len(h) }
func (h itemHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h itemHeap) Less(i, j int) bool {
	return h[i].val.Less(h[j].val)
}
func (h *itemHeap) Push(x interface{}) {
	*h = append(*h, x.(item))
}
func (h *itemHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
