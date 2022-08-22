package p0857minimumcosttohirekworkers

import (
	"container/heap"
	"math"
	"sort"
)

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	// Given a group of workers, the total amount paid will depend on the person
	// with the highest ratio of wage/quality.
	// So given a certain person as the highest wage/quality individual, the
	// optimal matching is the set of other people with lower wage/quality such
	// that the total quality is as small as possible.
	//
	// This gives us the solution - we sort workers by wage/quality ratio.
	// Then given each worker, we pair with the smallest total quality possible
	// from prior workers. This is the optimal group with that person as the
	// highest wage/quality individual.
	//
	type worker struct {
		ratio   float64
		quality int
	}
	n := len(quality)
	workers := make([]worker, n)
	for i := range quality {
		workers[i] = worker{float64(wage[i]) / float64(quality[i]), quality[i]}
	}
	sort.Slice(workers, func(i, j int) bool {
		return workers[i].ratio < workers[j].ratio
	})
	h := maxHeap{}
	var groupQuality int
	minCost := float64(math.MaxInt32)
	for i, w := range workers {
		groupQuality += w.quality
		if i >= k {
			// Remove largest quality in group
			groupQuality -= heap.Pop(&h).(int)
		}
		heap.Push(&h, w.quality)
		if i >= k-1 {
			minCost = math.Min(minCost, float64(groupQuality)*w.ratio)
		}
	}
	return minCost
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
