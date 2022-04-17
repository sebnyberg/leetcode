package p2244minimumroundstocompletealltasks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumRounds(t *testing.T) {
	for _, tc := range []struct {
		tasks []int
		want  int
	}{
		{[]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}, 4},
		{[]int{2, 3, 3}, -1},
		{[]int{2, 2, 2, 2}, 2},
		{[]int{2, 2, 2, 2, 2}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tasks), func(t *testing.T) {
			require.Equal(t, tc.want, minimumRounds(tc.tasks))
		})
	}
}

func minimumRounds(tasks []int) int {
	// The only invalid difficulty is 1. Else it's possible to do with 3s and 2s.
	// If evenly divisible by 3 => divide by 3
	// If mod 1 => divide by 3 and add 1
	// If mod 2 => divide by 3 and add 1
	m := make(map[int]int)
	for _, t := range tasks {
		m[t]++
	}
	var res int
	for _, v := range m {
		if v == 1 {
			return -1
		}
		if v%3 == 0 {
			res += v / 3
		} else {
			res += v/3 + 1
		}
	}
	return res
}
