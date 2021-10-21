package p1953maxnumberofweeksforwhichyoucanwork

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfWeeks(t *testing.T) {
	for _, tc := range []struct {
		milestones []int
		want       int64
	}{
		{[]int{9, 3, 6, 8, 2, 1}, 29},
		{[]int{1, 2, 3}, 6},
		{[]int{5, 2, 1}, 7},
	} {
		t.Run(fmt.Sprintf("%+v", tc.milestones), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfWeeks(tc.milestones))
		})
	}
}

func numberOfWeeks(milestones []int) int64 {
	// If the largest project is larger than the sum of all other projects, then
	// there is no way to finish all projects. Otherwise, it is possible.
	var maxVal, sum int
	for _, m := range milestones {
		sum += m
		if m > maxVal {
			maxVal = m
		}
	}
	if d := sum - maxVal; d < maxVal {
		return int64(d*2 + 1)
	}
	return int64(sum)
}
