package p2895minimumprocessingtime

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minProcessingTime(t *testing.T) {
	for i, tc := range []struct {
		processorTime []int
		tasks         []int
		want          int
	}{
		{[]int{10, 20}, []int{2, 3, 1, 2, 5, 8, 4, 3}, 23},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minProcessingTime(tc.processorTime, tc.tasks))
		})
	}
}

func minProcessingTime(processorTime []int, tasks []int) int {
	sort.Ints(processorTime)
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] > tasks[j]
	})
	var j int
	var res int
	for i := 0; i < len(tasks); i += 4 {
		res = max(res, processorTime[j]+tasks[i])
		j++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
