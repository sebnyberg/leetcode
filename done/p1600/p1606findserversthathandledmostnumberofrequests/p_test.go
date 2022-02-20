package p1606findserversthathandledmostnumberofrequests

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_busiestServers(t *testing.T) {
	for _, tc := range []struct {
		k       int
		arrival []int
		load    []int
		want    []int
	}{
		{7, []int{1, 3, 4, 5, 6, 11, 12, 13, 15, 19, 20, 21, 23, 25, 31, 32}, []int{9, 16, 14, 1, 5, 15, 6, 10, 1, 1, 7, 5, 11, 4, 4, 6}, []int{0}},
		{2, []int{1, 2, 3}, []int{1000000000, 1, 1000000000}, []int{1}},
		{3, []int{1, 2, 3, 4, 8, 9, 10}, []int{5, 2, 10, 3, 1, 2, 2}, []int{1}},
		{3, []int{1, 2, 3, 4, 5}, []int{5, 2, 3, 3, 3}, []int{1}},
		{3, []int{1, 2, 3, 4}, []int{1, 2, 1, 2}, []int{0}},
		{3, []int{1, 2, 3}, []int{10, 11, 12}, []int{0, 1, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, busiestServers(tc.k, tc.arrival, tc.load))
		})
	}
}

type server struct {
	idx     int
	heapIdx int
}

func busiestServers(k int, arrival []int, load []int) []int {
	requests := make([]int, k)
	servers := make([]*server, k)
	before := make(workerHeap, 0, k)
	after := make(workerHeap, k)
	for i := 0; i < k; i++ {
		servers[i] = &server{i, -1}
		after[i] = servers[i]
	}
	heap.Init(&after)

	jobs := make(jobHeap, 0, k)
	for i := range arrival {
		t := arrival[i]
		finish := t + load[i]

		// Workers that are finished are now available
		for len(jobs) > 0 && jobs[0].t <= t {
			j := heap.Pop(&jobs).(job)
			if j.idx < i%k {
				heap.Push(&before, servers[j.idx])
			} else {
				heap.Push(&after, servers[j.idx])
			}
		}

		var firstAvailable *server
		if len(after) > 0 {
			firstAvailable = heap.Pop(&after).(*server)
		} else if len(before) > 0 {
			firstAvailable = heap.Pop(&before).(*server)
		} else {
			continue
		}
		requests[firstAvailable.idx]++
		firstAvailable.heapIdx = -1
		heap.Push(&jobs, job{idx: firstAvailable.idx, t: finish})

		// Tricky part - the before/after boundary has moved
		// Move current index from after to before (if it's in the heap)
		if servers[i%k].heapIdx != -1 {
			heap.Remove(&after, servers[i%k].heapIdx)
			heap.Push(&before, servers[i%k])
		}

		if i%k == k-1 {
			// Flip before/after
			before, after = after, before
		}
	}
	var maxRequests int
	var res []int
	for i, count := range requests {
		if count > maxRequests {
			res = res[:0]
			maxRequests = count
		}
		if count == maxRequests {
			res = append(res, i)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type workerHeap []*server

func (h workerHeap) Len() int { return len(h) }
func (h workerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIdx = i
	h[j].heapIdx = j
}
func (h workerHeap) Less(i, j int) bool {
	return h[i].idx < h[j].idx
}
func (h *workerHeap) Push(x interface{}) {
	it := x.(*server)
	n := len(*h)
	it.heapIdx = n
	*h = append(*h, it)
}
func (h *workerHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

type job struct {
	idx int
	t   int
}

type jobHeap []job

func (h jobHeap) Len() int { return len(h) }
func (h jobHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h jobHeap) Less(i, j int) bool {
	return h[i].t < h[j].t
}
func (h *jobHeap) Push(x interface{}) {
	*h = append(*h, x.(job))
}
func (h *jobHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
