package p2071maxnumberoftasksyoucanassign

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxTaskAssign(t *testing.T) {
	for _, tc := range []struct {
		tasks    []int
		workers  []int
		pills    int
		strength int
		want     int
	}{
		{[]int{5, 9, 8, 5, 9}, []int{1, 6, 4, 2, 6}, 1, 5, 3},
		{[]int{3, 2, 1}, []int{0, 3, 3}, 1, 1, 3},
		{[]int{5, 4}, []int{0, 0, 0}, 1, 5, 1},
		{[]int{10, 15, 30}, []int{0, 10, 10, 10, 10}, 3, 10, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tasks), func(t *testing.T) {
			require.Equal(t, tc.want, maxTaskAssign(tc.tasks, tc.workers, tc.pills, tc.strength))
		})
	}
}

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	n := len(tasks)

	// Sort workers descending
	sort.Slice(workers, func(i, j int) bool {
		return workers[i] > workers[j]
	})
	// Sort work descending
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] > tasks[j]
	})

	maxPossible := min(n, len(workers))
	checkCanCompleteTasks := func(x int) bool {
		if x > maxPossible {
			return false
		}
		workersCpy := make([]int, len(workers))
		copy(workersCpy, workers)
		workers := workersCpy
		npills := pills
		for i := 0; i < x; i++ {
			t := tasks[n-x+i]
			if workers[0] >= t {
				workers = workers[1:]
				continue
			}
			// Check if we can give a pill to a worker instead...
			if npills == 0 {
				return false
			}
			// Add pill to weakest worker who has enough strength to do the task
			firstTooWeakIdx := sort.Search(len(workers), func(j int) bool {
				return workers[j]+strength < t
			})
			if firstTooWeakIdx == 0 { // no matching worker
				return false
			}
			copy(workers[firstTooWeakIdx-1:], workers[firstTooWeakIdx:])
			npills--
		}
		return true
	}

	l, r := 0, maxPossible+1
	for l < r {
		mid := (l + r) / 2
		if checkCanCompleteTasks(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l - 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
