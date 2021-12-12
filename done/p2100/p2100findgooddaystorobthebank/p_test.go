package p2100findgooddaystorobthebank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_goodDaysToRobBank(t *testing.T) {
	for _, tc := range []struct {
		security []int
		time     int
		want     []int
	}{
		{[]int{5, 3, 3, 3, 5, 6, 2}, 2, []int{2, 3}},
		{[]int{1, 1, 1, 1, 1}, 0, []int{0, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6}, 2, []int{}},
		{[]int{1}, 5, []int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.security), func(t *testing.T) {
			require.Equal(t, tc.want, goodDaysToRobBank(tc.security, tc.time))
		})
	}
}

func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	valid := make([]bool, n)
	count := 1
	for i := range security {
		if i > 0 {
			if security[i-1] >= security[i] {
				count++
			} else {
				count = 1
			}
		}
		if count > time {
			valid[i] = true
		}
	}

	count = 1
	for i := n - 1; i >= 0; i-- {
		if i < n-1 {
			if security[i] <= security[i+1] {
				count++
			} else {
				count = 1
			}
		}
		valid[i] = valid[i] && count > time
	}

	res := make([]int, 0)
	for i := range valid {
		if valid[i] {
			res = append(res, i)
		}
	}

	return res
}
