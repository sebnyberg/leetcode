package p2274maximumconsecutivefloorswithoutspecialfloors

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxConsecutive(t *testing.T) {
	for _, tc := range []struct {
		bottom  int
		top     int
		special []int
		want    int
	}{
		{2, 9, []int{4, 6}, 3},
		{6, 8, []int{7, 6, 8}, 0},
		{3, 15, []int{7, 9, 13}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.bottom), func(t *testing.T) {
			require.Equal(t, tc.want, maxConsecutive(tc.bottom, tc.top, tc.special))
		})
	}
}

func maxConsecutive(bottom int, top int, special []int) int {
	special = append(special, top+1, bottom-1)
	sort.Ints(special)
	var start, end int
	for i := range special {
		if special[i] < bottom-1 {
			start++
			continue
		}
		if special[i] > top {
			end = i + 1
			break
		}
	}
	special = special[start:end]
	var maxRes int
	for i := 1; i < len(special); i++ {
		maxRes = max(maxRes, (special[i]-special[i-1])-1)
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
