package p2028findmissingobservations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_missingRolls(t *testing.T) {
	for _, tc := range []struct {
		rolls   []int
		mean, n int
		want    []int
	}{
		{[]int{3, 2, 4, 3}, 4, 2, []int{6, 6}},
		{[]int{1, 5, 6}, 3, 4, []int{2, 3, 2, 2}},
		{[]int{1, 2, 3, 4}, 6, 4, []int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rolls), func(t *testing.T) {
			res := missingRolls(tc.rolls, tc.mean, tc.n)
			if len(res) == 0 {
				require.Equal(t, tc.want, res)
				return
			}
			var sum int
			for _, roll := range tc.rolls {
				sum += roll
			}
			for _, resItem := range res {
				sum += resItem
			}
			require.Equal(t, tc.mean, sum/(len(tc.rolls)+tc.n))
		})
	}
}

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	var sum int
	for _, roll := range rolls {
		sum += roll
	}
	wantSum := mean * (m + n)
	missing := wantSum - sum
	if missing < n || missing > 6*n {
		return []int{}
	}
	extra := missing - n
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
		if extra > 0 {
			toAdd := min(extra, 5)
			res[i] += toAdd
			extra -= toAdd
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
