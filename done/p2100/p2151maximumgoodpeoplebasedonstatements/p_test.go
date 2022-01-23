package p2151maximumgoodpeoplebasedonstatements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumGood(t *testing.T) {
	for _, tc := range []struct {
		statements [][]int
		want       int
	}{
		{[][]int{{2, 1, 2}, {1, 2, 2}, {2, 0, 2}}, 2},
		{[][]int{{2, 0}, {0, 2}}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.statements), func(t *testing.T) {
			require.Equal(t, tc.want, maximumGood(tc.statements))
		})
	}
}

const (
	statusBad  = 0
	statusGood = 1
	statusNone = 2
)

type edge struct {
	from   int
	to     int
	status int8
}

func maximumGood(statements [][]int) int {
	n := len(statements)
	status := make([]int8, n)
	res := dfs(status, statements, 0, n)
	return res
}

// dfs returns true if all nodes were covered
func dfs(status []int8, statements [][]int, idx, n int) int {
	if idx == n {
		// Check if valid
		var count int
		for i := 0; i < n; i++ {
			if status[i] == statusBad { // This person is not trustworthy
				continue
			}
			count++
			// Validate that it's possible that this person is good
			// If this person is good, then his opinions must match the config
			for j := 0; j < n; j++ {
				st := int8(statements[i][j])
				if st != statusNone && st != status[j] {
					return -1
				}
			}
		}
		return count
	}
	// Set current status to good
	status[idx] = statusGood
	res := dfs(status, statements, idx+1, n)
	// Set to bad
	status[idx] = statusBad
	if other := dfs(status, statements, idx+1, n); other > res {
		return other
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
