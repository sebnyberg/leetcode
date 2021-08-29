package p1986minnumberofworksessionsto

import (
	"fmt"
	"math"
	"testing"
)

func Test_minSessions(t *testing.T) {
	type args struct {
		tasks       []int
		sessionTime int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{[]int{1, 2, 3}, 3}, 2},
		{args{[]int{3, 1, 3, 1, 1}, 8}, 2},
		{args{[]int{1, 2, 3, 4, 5}, 15}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.args), func(t *testing.T) {
			if got := minSessions(tt.args.tasks, tt.args.sessionTime); got != tt.want {
				t.Errorf("minSessions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func minSessions(tasks []int, sessionTime int) int {
	// This is a knapsack-style problem
	// The goal is to pack tasks in such a way that the total number of knapsacks
	// is minimized.
	// Idea:
	// 1. Store picked sessions in a bitmap
	// 2. For each number of tasks, i = 1, ..., len(tasks)
	// 3. For each task index, try adding that task to each bitmap which did not
	//    include the task.
	sessionOrdinal := 100
	dp := make(map[int]int)
	dp[0] = 0
	for n := 1; n <= len(tasks); n++ {
		nextDP := make(map[int]int)
		for prevBM, val := range dp {
			for i, taskLen := range tasks {
				if prevBM&(1<<i) > 0 {
					// bitmap already contains this task
					continue
				}
				// combine previous task(s) with current task, check if smaller
				// than previous best result
				bm := prevBM | (1 << i)
				if _, exists := nextDP[bm]; !exists {
					nextDP[bm] = math.MaxInt32
				}
				res := val + taskLen
				if res%sessionOrdinal > sessionTime {
					res = sessionOrdinal*((val/sessionOrdinal)+1) + taskLen
				} else if res%sessionOrdinal == sessionTime {
					res = sessionOrdinal * ((val / sessionOrdinal) + 1)
				}
				nextDP[bm] = min(nextDP[bm], res)
			}
		}
		dp = nextDP
	}
	finalBM := 1<<len(tasks) - 1
	res := dp[finalBM]
	nsessions := res / sessionOrdinal
	if res%sessionOrdinal != 0 {
		nsessions++
	}
	return nsessions
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
