package p1383maxperfofteam

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxPerformance(t *testing.T) {
	for _, tc := range []struct {
		n          int
		speed      []int
		efficiency []int
		k          int
		want       int
	}{
		{3, []int{2, 8, 2}, []int{2, 7, 1}, 2, 56},
		{6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 2, 60},
		{6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 3, 68},
		{6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 4, 72},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, maxPerformance(tc.n, tc.speed, tc.efficiency, tc.k))
		})
	}
}

type engineer struct {
	speed      int
	efficiency int
}

type speedHeap []int

func (h speedHeap) Len() int      { return len(h) }
func (h speedHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *speedHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	(*h) = (*h)[:n-1]
	return it
}
func (h *speedHeap) Push(x interface{}) { (*h) = append((*h), x.(int)) }
func (h speedHeap) Less(i, j int) bool  { return h[i] < h[j] }

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	// 1. sort by efficiency
	// 2. pick top speediest engineers
	// We do not care about order, only the max value
	engineers := make([]engineer, n)
	for i := range speed {
		engineers[i] = engineer{
			speed:      speed[i],
			efficiency: efficiency[i],
		}
	}
	sort.Slice(engineers, func(i, j int) bool {
		return engineers[i].efficiency > engineers[j].efficiency
	})

	maxSpeeds := make(speedHeap, 0)
	heap.Init(&maxSpeeds)
	groupSpeed := 0
	var maxResult int
	for i := 0; i < n; i++ {
		groupSpeed += engineers[i].speed
		maxResult = max(maxResult, groupSpeed*engineers[i].efficiency)
		heap.Push(&maxSpeeds, engineers[i].speed)
		if len(maxSpeeds) >= k {
			groupSpeed -= heap.Pop(&maxSpeeds).(int)
		}
	}

	return maxResult % 1000000007
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
