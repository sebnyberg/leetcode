package p1856maximumsubarrayminproduct

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSumMinProduct(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 3, 1, 2}, 18},
		{[]int{1, 2, 3, 2}, 14},
		{[]int{3, 1, 5, 6, 4}, 60},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxSumMinProduct(tc.nums))
		})
	}

}

const mod = 1000000007

func maxSumMinProduct(nums []int) int {
	h := make(MaxHeap, len(nums))
	for i, n := range nums {
		h[i] = &Item{val: n, idx: i}
	}
	heap.Init(&h)
	seen := make([]bool, len(nums))
	dsu := NewDSU(nums)
	for len(h) > 0 {
		it := heap.Pop(&h).(*Item)
		seen[it.idx] = true
		for _, nei := range []int{it.idx - 1, it.idx + 1} {
			if nei < 0 || nei >= len(nums) {
				continue
			}
			if seen[nei] {
				dsu.union(it.idx, nei)
			}
		}
	}
	return dsu.maxMinProduct % mod
}

type Item struct {
	val int
	idx int
}

type MaxHeap []*Item

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i].val > h[j].val
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

type DSU struct {
	parent        []int
	sums          []int
	minVals       []int
	maxMinProduct int
}

func NewDSU(nums []int) *DSU {
	dsu := &DSU{
		parent:  make([]int, len(nums)),
		sums:    make([]int, len(nums)),
		minVals: make([]int, len(nums)),
	}
	for i, n := range nums {
		dsu.parent[i] = i
		dsu.sums[i] = n
		dsu.minVals[i] = n
		dsu.maxMinProduct = max(dsu.maxMinProduct, n*n)
	}
	return dsu
}

func (d *DSU) find(a int) int {
	if d.parent[a] == a {
		return a
	}
	root := d.find(d.parent[a])
	d.parent[a] = root
	return root
}

func (d *DSU) union(a, b int) {
	a = d.find(a)
	b = d.find(b)
	if a != b {
		d.parent[b] = a
		d.sums[a] += d.sums[b]
		d.minVals[a] = min(d.minVals[b], d.minVals[a])
		d.maxMinProduct = max(d.maxMinProduct, d.sums[a]*d.minVals[a])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
