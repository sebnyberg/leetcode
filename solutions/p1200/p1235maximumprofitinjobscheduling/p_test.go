package p1235maximumprofitinjobscheduling

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_jobScheduling(t *testing.T) {
	for _, tc := range []struct {
		startTime, endTime, profit []int
		want                       int
	}{
		{[]int{4, 2, 4, 8, 2}, []int{5, 5, 5, 10, 8}, []int{1, 2, 8, 10, 4}, 18},
		{[]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}, 120},
		{[]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}, 150},
		{[]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.startTime), func(t *testing.T) {
			require.Equal(t, tc.want, jobScheduling(tc.startTime, tc.endTime, tc.profit))
		})
	}
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	type task struct {
		start, end, profit, maxProfit int
	}
	n := len(startTime)
	tasks := make([]task, n)
	for i := range startTime {
		tasks[i] = task{startTime[i], endTime[i], profit[i], profit[i]}
	}

	// Sort by end-time
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].end < tasks[j].end
	})
	tasks = append(tasks, task{0, 1e9 + 1, 0, 0}) // sentinel

	// Visit tasks one by one
	for i := 1; i < len(tasks)-1; i++ {
		// Find first task with an end time <= start time for current
		// This is done by binary searching for the first time greater
		gtIdx := sort.Search(i, func(j int) bool {
			return tasks[j].end > tasks[i].start
		})
		for j := gtIdx - 1; j >= 0 && tasks[j].end == tasks[gtIdx-1].end; j-- {
			tasks[i].maxProfit = max(tasks[i].maxProfit, tasks[i].profit+tasks[j].maxProfit)
		}
		tasks[i].maxProfit = max(tasks[i].maxProfit, tasks[i-1].maxProfit)
	}
	return tasks[len(tasks)-2].maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
