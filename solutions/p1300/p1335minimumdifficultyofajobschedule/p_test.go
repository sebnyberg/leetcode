package p_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDifficulty(t *testing.T) {
	for i, tc := range []struct {
		jobDifficulty []int
		d             int
		want          int
	}{
		// {[]int{6, 5, 4, 3, 2, 1}, 2, 7},
		// {[]int{9, 9, 9}, 4, -1},
		{[]int{1, 1, 1}, 3, 3},
		// {[]int{7,1,7,1,7,1}, 3, 15},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minDifficulty(tc.jobDifficulty, tc.d))
		})
	}
}

func minDifficulty(jobDifficulty []int, d int) int {
	if len(jobDifficulty) < d {
		return -1
	}
	var mem [301][11]int
	for i := range mem {
		for j := range mem[i] {
			mem[i][j] = math.MaxInt32
		}
	}
	n := len(jobDifficulty)
	res := dfs(&mem, jobDifficulty, n, d)
	return res
}

func dfs(mem *[301][11]int, jobs []int, j, d int) int {
	if mem[j][d] < math.MaxInt32 {
		return mem[j][d]
	}
	if j == 1 {
		if d == 1 {
			return jobs[0]
		}
		return math.MaxInt32
	}
	if d == 1 {
		return max(jobs[j-1], dfs(mem, jobs, j-1, 1))
	}
	res := math.MaxInt32
	var maxVal int
	for jj := j - 1; jj >= 1; jj-- {
		maxVal = max(maxVal, jobs[jj])
		res = min(res, maxVal+dfs(mem, jobs, jj, d-1))
	}
	mem[j][d] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
