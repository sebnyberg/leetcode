package p0502ipo

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMaximizedCapital(t *testing.T) {
	for _, tc := range []struct {
		k       int
		w       int
		profits []int
		capital []int
		want    int
	}{
		{2, 0, []int{1, 2, 3}, []int{0, 1, 1}, 4},
		{3, 0, []int{1, 2, 3}, []int{0, 1, 2}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, findMaximizedCapital(tc.k, tc.w, tc.profits, tc.capital))
		})
	}
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	// We want to pick the highest profit project such that its capital required
	// is smaller than or equal to the current capital
	//
	// Idea: add all projects which have <= capital in requirement to a max-heap
	// Then pick the maximum item from the heap each time, updating the heap with
	// new items which are made available by the new capital
	n := len(profits)
	projects := make([]*project, n)
	for i, p := range profits {
		projects[i] = &project{profit: p, capital: capital[i]}
	}
	h := make(projectHeap, 0, n)
	copy(h, projects)
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})
	var j int
	for ; j < n && projects[j].capital <= w; j++ {
		h = append(h, projects[j])
	}
	heap.Init(&h)
	for k > 0 && len(h) > 0 {
		// Get most profitable project
		x := heap.Pop(&h).(*project)
		w += x.profit
		for ; j < n && projects[j].capital <= w; j++ {
			heap.Push(&h, projects[j])
		}
		k--
	}
	return w
}

type project struct {
	profit  int
	capital int
}

type projectHeap []*project

func (h projectHeap) Len() int { return len(h) }
func (h projectHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]

}
func (h projectHeap) Less(i, j int) bool {
	return h[i].profit > h[j].profit
}
func (h *projectHeap) Push(x interface{}) {
	*h = append(*h, x.(*project))
}
func (h *projectHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
