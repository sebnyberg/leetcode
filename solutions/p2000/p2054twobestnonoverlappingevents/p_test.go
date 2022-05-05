package p2054twobestnonoverlappingevents

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxTwoEvents(t *testing.T) {
	for _, tc := range []struct {
		events [][]int
		want   int
	}{
		{[][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}, 4},
		{[][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}, 5},
		{[][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}, 8},
	} {
		t.Run(fmt.Sprintf("%+v", tc.events), func(t *testing.T) {
			require.Equal(t, tc.want, maxTwoEvents(tc.events))
		})
	}
}

func maxTwoEvents(events [][]int) int {
	// The max value from attending one event is just the maximum value of any
	// event.

	// The max value of attenting two events is the max value from a position +
	// the max value prior to that position.

	// When evaluating a start-date, update maximum of two events to be maxPrior
	// + value of current event

	// When evaluating an end-date, update max prior to be max(maxPrior, current)
	h := make(EventHeap, 0, len(events)*2)
	for _, ev := range events {
		h = append(h, EventBoundary{ev[0], ev[2]})
		h = append(h, EventBoundary{ev[1], -ev[2]})
	}
	heap.Init(&h)
	var maxPrior int
	var maxTotal int
	for len(h) > 0 {
		t := h[0].time
		var maxPriorToCurrent int
		for len(h) > 0 && h[0].time == t {
			e := heap.Pop(&h).(EventBoundary)
			if e.val < 0 { // end of an event
				maxPriorToCurrent = max(maxPriorToCurrent, -e.val)
			} else { // start of event
				maxTotal = max(maxTotal, maxPrior+e.val)
			}
		}
		maxPrior = max(maxPrior, maxPriorToCurrent)
	}
	return maxTotal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type EventBoundary struct {
	time int
	val  int
}

type EventHeap []EventBoundary

func (h EventHeap) Len() int { return len(h) }
func (h EventHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h EventHeap) Less(i, j int) bool {
	return h[i].time < h[j].time
}
func (h *EventHeap) Push(x interface{}) {
	*h = append(*h, x.(EventBoundary))
}
func (h *EventHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
