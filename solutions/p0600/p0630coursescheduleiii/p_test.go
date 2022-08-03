package p0630coursescheduleiii

import (
	"container/heap"
	"fmt"
	"github.com/sebnyberg/leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_scheduleCourse(t *testing.T) {
	for _, tc := range []struct {
		courses [][]int
		want    int
	}{
		{leetcode.ParseMatrix("[[100,200],[200,1300],[1000,1250],[2000,3200]]"), 3},
		{leetcode.ParseMatrix("[[1,2]]"), 1},
		{leetcode.ParseMatrix("[[3,2],[4,3]]"), 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.courses), func(t *testing.T) {
			require.Equal(t, tc.want, scheduleCourse(tc.courses))
		})
	}
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	h := maxHeap{}
	var t int
	for _, c := range courses {
		t += c[0]
		heap.Push(&h, c[0])
		if t > c[1] {
			t -= heap.Pop(&h).(int)
		}
	}
	return len(h)
}

type maxHeap []int

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
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
