package p1353maximumnumberofeventsthatcanbeattended

import (
	"container/heap"
	"sort"
)

func maxEvents(events [][]int) int {
	// For each point in time, we want to attend the event that is closing
	// soonest.

	h := minHeap{}
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	var res int
	t := events[0][0]
	var i int
	for i < len(events) || len(h) > 0 {
		for i < len(events) && events[i][0] <= t {
			heap.Push(&h, events[i][1])
			i++
		}
		// For the time t - what is the event that will end soonest?
		//
		// Remove out of date events
		for len(h) > 0 && h[0] < t {
			heap.Pop(&h)
		}

		// If there are no ongoing events, change t to time of next event
		if len(h) == 0 {
			if i >= len(events) {
				break
			}
			t = events[i][0]
			continue
		}

		// Pick an event
		heap.Pop(&h)
		res++
		t++
	}

	return res
}

type minHeap []int

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *minHeap) Push(x interface{}) {
	el := x.(int)
	*h = append(*h, el)
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	*h = (*h)[:n-1]
	return el
}
