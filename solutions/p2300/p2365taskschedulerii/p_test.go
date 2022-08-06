package p2365taskschedulerii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_taskSchedulerII(t *testing.T) {
	for _, tc := range []struct {
		tasks []int
		space int
		want  int64
	}{
		{[]int{8}, 1, 1},
		{[]int{1, 2, 3}, 1, 3},
		{[]int{5, 8, 8, 5}, 2, 6},
		{[]int{1, 2, 1, 2, 3, 1}, 3, 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tasks), func(t *testing.T) {
			require.Equal(t, tc.want, taskSchedulerII(tc.tasks, tc.space))
		})
	}
}

// https://www.github.com/sebnyberg/leetcode
func taskSchedulerII(tasks []int, space int) int64 {
	m := make(map[int]int64)
	var t int64
	for _, task := range tasks {
		if _, exists := m[task]; !exists {
			t++
		} else {
			t = max(t+1, m[task])
		}
		m[task] = t + int64(space) + 1
	}
	return t
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
