package p0630courseschedule3

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_scheduleCourse(t *testing.T) {
	for _, tc := range []struct {
		courses [][]int
		want    int
	}{
		{[][]int{{5, 5}, {4, 6}, {2, 6}}, 2},
		{[][]int{{5, 15}, {3, 19}, {6, 7}, {2, 10}, {5, 16}, {8, 14}, {10, 11}, {2, 19}}, 5},
		{[][]int{{100, 200}, {100, 200}}, 2},
		{[][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}, 3},
		{[][]int{{1, 2}}, 1},
		{[][]int{{3, 2}, {4, 3}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.courses), func(t *testing.T) {
			require.Equal(t, tc.want, scheduleCourse(tc.courses))
		})
	}
}

func scheduleCourse(courses [][]int) int {
	// Sort slices by last day (low -> high)
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	// Add courses to duration heap
	h := MaxHeap{}
	t := 0
	for _, course := range courses {
		duration, end := course[0], course[1]
		// If course can be taken
		if t+duration <= end {
			// add it to the max heap
			heap.Push(&h, duration)
			t += duration
			continue
		}
		// If course cannot be taken with the current time,
		// and the current course has shorter duration than a taken course
		// remove taken course and use this one instead (optimal)
		if len(h) > 0 && h[0] > duration {
			t += duration - heap.Pop(&h).(int)
			heap.Push(&h, duration)
		}
	}
	return len(h)
}

type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
