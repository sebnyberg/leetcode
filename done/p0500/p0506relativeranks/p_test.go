package p0506relativeranks

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRelativeRanks(t *testing.T) {
	for _, tc := range []struct {
		score []int
		want  []string
	}{
		{[]int{5, 4, 3, 2, 1}, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
		{[]int{10, 3, 8, 9, 4}, []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.score), func(t *testing.T) {
			require.Equal(t, tc.want, findRelativeRanks(tc.score))
		})
	}
}

func findRelativeRanks(score []int) []string {
	sorted := make([]int, len(score))
	copy(sorted, score)
	sort.Ints(sorted)
	rank := make(map[int]int)
	for i, num := range sorted {
		rank[num] = len(score) - i
	}
	res := make([]string, len(score))
	for i, num := range score {
		switch rank[num] {
		case 1:
			res[i] = "Gold Medal"
		case 2:
			res[i] = "Silver Medal"
		case 3:
			res[i] = "Bronze Medal"
		default:
			res[i] = fmt.Sprint(rank[num])
		}
	}
	return res
}
