package p1834singlethreadedcpu

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getOrder(t *testing.T) {
	for _, tc := range []struct {
		tasks [][]int
		want  []int
	}{
		{[][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}}, []int{6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7}},
		{[][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}, []int{0, 2, 3, 1}},
		{[][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}, []int{4, 3, 2, 0, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tasks), func(t *testing.T) {
			require.Equal(t, tc.want, getOrder(tc.tasks))
		})
	}
}

func getOrder(tasks [][]int) []int {
	// sort tasks by time
	sortedTasks := make([]*cpuTask, len(tasks))
	for i, t := range tasks {
		sortedTasks[i] = &cpuTask{
			idx:            i,
			enqueueTime:    t[0],
			processingTime: t[1],
		}
	}
	sort.Slice(sortedTasks, func(i, j int) bool {
		if sortedTasks[i].enqueueTime == sortedTasks[j].enqueueTime {
			return sortedTasks[i].processingTime < sortedTasks[j].processingTime
		}
		return sortedTasks[i].enqueueTime < sortedTasks[j].enqueueTime
	})
	h := make(MinHeap, 0)
	res := make([]int, 0, len(tasks))
	t := 0
	var cur *cpuTask
	for len(h) > 0 || len(sortedTasks) > 0 {
		if len(h) == 0 {
			// pick first task in list
			cur = sortedTasks[0]
			sortedTasks = sortedTasks[1:]
		} else {
			// pick shortest processing time from heap
			cur = heap.Pop(&h).(*cpuTask)
		}
		if cur.enqueueTime > t {
			t = cur.enqueueTime
		}
		res = append(res, cur.idx)
		// add all from current time until end of processing time
		t += cur.processingTime
		for len(sortedTasks) > 0 && sortedTasks[0].enqueueTime <= t {
			heap.Push(&h, sortedTasks[0])
			sortedTasks = sortedTasks[1:]
		}
	}
	return res
}

type cpuTask struct {
	idx            int
	enqueueTime    int
	processingTime int
}

type MinHeap []*cpuTask

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Less(i, j int) bool {
	if h[i].processingTime == h[j].processingTime {
		return h[i].idx < h[j].idx
	}
	return h[i].processingTime < h[j].processingTime
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*cpuTask))
}
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
