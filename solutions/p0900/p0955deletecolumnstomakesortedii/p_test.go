package p0955deletecolumnstomakesortedii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDeletionSize(t *testing.T) {
	for i, tc := range []struct {
		strs []string
		want int
	}{
		{[]string{"vdy", "vei", "zvc", "zld"}, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minDeletionSize(tc.strs))
		})
	}
}

func minDeletionSize(strs []string) int {
	// Unless I'm missing something, this is just a greedy problem
	// Ensure that each pair is sorted and move on

	n := len(strs[0])
	m := len(strs)
	skip := make([]bool, n)
	for i := 1; i < m; i++ {
		// Check this pair
		for j := 0; j < n; j++ {
			if skip[j] {
				continue
			}
			if strs[i-1][j] > strs[i][j] {
				skip[j] = true
				i = 0
				break
			} else if strs[i-1][j] < strs[i][j] {
				break
			}
		}
	}
	var res int
	for _, v := range skip {
		if v {
			res++
		}
	}
	return res
}
