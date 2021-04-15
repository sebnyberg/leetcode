package p1785minelementtoaddtoformagivensum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minElements(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		limit int
		goal  int
		want  int
	}{
		{[]int{1, 2, 3}, 300, 6, 0},
		{[]int{1, -1, 1}, 3, -4, 2},
		{[]int{1, -10, 9, 1}, 100, 0, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minElements(tc.nums, tc.limit, tc.goal))
		})
	}
}

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum == goal {
		return 0
	}
	return ((abs(sum-goal) - 1) / limit) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
