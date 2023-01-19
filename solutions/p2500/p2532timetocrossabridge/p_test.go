package p2532timetocrossabridge

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_findCrossingTime(t *testing.T) {
	for i, tc := range []struct {
		n    int
		k    int
		time [][]int
		want int
	}{
		{1, 3, leetcode.ParseMatrix("[[1,1,2,1],[1,1,3,1],[1,1,4,1]]"), 6},
		{3, 2, leetcode.ParseMatrix("[[1,9,1,8],[10,10,10,10]]"), 50},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, findCrossingTime(tc.n, tc.k, tc.time))
		})
	}
}

func findCrossingTime(n int, k int, time [][]int) int {
	// Everyone starts on the left side
	// As soon as someone crosses, they start picking up or leaving a box
	// Returning workers have priority over leaving workers
	// There is no reason for a worker to leave if there are no more boxes on
	// the other side.
	var busy workerHeap
	var waiting workerHeap
	for i, t := range time {
		st := workerState{
			idx:     i,
			lr:      t[0],
			pickOld: t[1],
			rl:      t[2],
			putNew:  t[3],
			leaving: true,
			readyAt: 0,
		}
		waiting = append(waiting, st)
	}
	heap.Init(&waiting)
	var t int
	for n > 0 || len(waiting) > 0 || len(busy) > 0 {
		if len(waiting) == 0 && busy[0].readyAt > t {
			// Fast-forward time if nobody is ready to cross the bridge
			t = busy[0].readyAt
		}
		for len(busy) > 0 && busy[0].readyAt <= t {
			x := heap.Pop(&busy).(workerState)
			x.readyAt = 0
			heap.Push(&waiting, x)
		}
		x := heap.Pop(&waiting).(workerState)
		if x.leaving && n == 0 {
			// Don't leave if there are no boxes on the other side
			continue
		}
		if x.leaving {
			// Worker is leaving to pick up a box
			n--
			x.readyAt = t + x.lr + x.pickOld
			x.leaving = false
			t += x.lr
			heap.Push(&busy, x)
		} else {
			// Worker is coming back
			x.readyAt = t + x.rl + x.putNew
			x.leaving = true
			t += x.rl
			if n != 0 {
				heap.Push(&busy, x)
			}
		}
	}
	return t
}

type workerState struct {
	idx     int
	lr      int
	rl      int
	pickOld int
	putNew  int
	leaving bool
	readyAt int
}

type workerHeap []workerState

func (h workerHeap) Len() int { return len(h) }
func (h workerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h workerHeap) Less(i, j int) bool {
	a := h[i]
	b := h[j]
	if a.readyAt != b.readyAt {
		return a.readyAt < b.readyAt
	}
	if a.leaving != b.leaving {
		return !a.leaving
	}
	aa := a.lr + a.rl
	bb := b.lr + b.rl
	if aa != bb {
		return aa > bb
	}
	return a.idx > b.idx
}
func (h *workerHeap) Push(x interface{}) {
	el := x.(workerState)
	*h = append(*h, el)
}
func (h *workerHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	*h = (*h)[:n-1]
	return el
}
