package p0826mostprofitassigningwork

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProfitAssignment(t *testing.T) {
	for _, tc := range []struct {
		difficulty []int
		profit     []int
		worker     []int
		want       int
	}{
		{[]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}, 100},
		{[]int{85, 47, 57}, []int{24, 66, 99}, []int{40, 25, 25}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.difficulty), func(t *testing.T) {
			require.Equal(t, tc.want, maxProfitAssignment(tc.difficulty, tc.profit, tc.worker))
		})
	}
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	// Remove jobs which have higher difficulty and lower profit
	n := len(difficulty)
	type jobItem struct {
		difficulty int
		profit     int
	}
	jobs := make([]jobItem, n)
	for i := range difficulty {
		jobs[i] = jobItem{difficulty[i], profit[i]}
	}
	jobs = append(jobs, jobItem{0, 0}) // sentinel
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].difficulty < jobs[j].difficulty
	})
	var j int
	for i := 1; i < len(jobs); i++ {
		if jobs[i].profit < jobs[j].profit {
			continue
		}
		j++
		jobs[j] = jobs[i]
	}
	jobs = jobs[:j+1]
	// Now the highest difficulty job is also the most profitable, so we can
	// search for each worker to get the result
	var res int
	for _, w := range worker {
		i := sort.Search(len(jobs), func(i int) bool {
			return jobs[i].difficulty > w
		})
		res += jobs[i-1].profit
	}
	return res
}
