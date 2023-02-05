package p2554maximumnumberofintegerstochoosefromarangei

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxCount(t *testing.T) {
	for i, tc := range []struct {
		banned []int
		n      int
		maxSum int
		want   int
	}{
		{[]int{1, 5, 6}, 5, 6, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 8, 1, 0},
		{[]int{11}, 7, 50, 7},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxCount(tc.banned, tc.n, tc.maxSum))
		})
	}
}

func maxCount(banned []int, n int, maxSum int) int {
	m := make(map[int]bool)
	for _, x := range banned {
		m[x] = true
	}
	var sum int
	var res int
	for k := 1; k <= n && sum+k <= maxSum; k++ {
		if m[k] {
			continue
		}
		sum += k
		res++
	}
	return res
}
