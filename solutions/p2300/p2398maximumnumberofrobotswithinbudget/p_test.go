package p2398maximumnumberofrobotswithinbudget

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumRobots(t *testing.T) {
	for _, tc := range []struct {
		chargeTimes  []int
		runningCosts []int
		budget       int64
		want         int
	}{
		{
			[]int{11, 12, 74, 67, 37, 87, 42, 34, 18, 90, 36, 28, 34, 20},
			[]int{18, 98, 2, 84, 7, 57, 54, 65, 59, 91, 7, 23, 94, 20},
			937,
			4,
		},
		{[]int{3, 6, 1, 3, 4}, []int{2, 1, 3, 4, 5}, 25, 3},
		{[]int{11, 12, 19}, []int{10, 8, 7}, 19, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.chargeTimes), func(t *testing.T) {
			require.Equal(t, tc.want, maximumRobots(tc.chargeTimes, tc.runningCosts, tc.budget))
		})
	}
}

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	// It's consecutive robots... I thought it was any set of robots.
	// This is much easier than I thought
	// Keep a heap of maximum charge times and sum of running costs.
	// For each element, add it to the heap and sum of running costs.
	// While the total sum is too large, remove from the heap
	var i int
	var sum int
	h := robotHeap{}
	n := len(chargeTimes)
	robots := make([]*robot, n)
	for i := range chargeTimes {
		robots[i] = &robot{0, chargeTimes[i]}
	}
	var res int
	for j := range chargeTimes {
		heap.Push(&h, robots[j])
		sum += runningCosts[j]
		for i <= j && int64(sum*(j-i+1)+h[0].chargeTime) > budget {
			sum -= runningCosts[i]
			heap.Remove(&h, robots[i].idx)
			i++
		}
		res = max(res, j-i+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type robot struct {
	idx        int
	chargeTime int
}

type robotHeap []*robot

func (h robotHeap) Len() int { return len(h) }
func (h robotHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}
func (h robotHeap) Less(i, j int) bool {
	return h[i].chargeTime > h[j].chargeTime
}
func (h *robotHeap) Push(x interface{}) {
	r := x.(*robot)
	r.idx = len(*h)
	*h = append(*h, r)
}
func (h *robotHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
