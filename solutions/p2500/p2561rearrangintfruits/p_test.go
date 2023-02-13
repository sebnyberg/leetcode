package p2561rearrangintfruits

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCost(t *testing.T) {
	for i, tc := range []struct {
		basket1 []int
		basket2 []int
		want    int64
	}{
		{
			[]int{84, 80, 43, 8, 80, 88, 43, 14, 100, 88},
			[]int{32, 32, 42, 68, 68, 100, 42, 84, 14, 8},
			48,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minCost(tc.basket1, tc.basket2))
		})
	}
}

func minCost(basket1 []int, basket2 []int) int64 {
	count := make(map[int]int)
	minVal := math.MaxInt32
	for i := range basket1 {
		count[basket1[i]]++
		count[basket2[i]]--
		minVal = min(minVal, basket1[i])
		minVal = min(minVal, basket2[i])
	}
	var nswaps int
	for _, c := range count {
		if c%2 != 0 {
			return -1
		}
		if c > 0 {
			nswaps += c / 2
		}
	}
	out1 := make([]int, 0, nswaps)
	out2 := make([]int, 0, nswaps)
	for fruit, cnt := range count {
		if cnt > 0 {
			for k := 0; k < cnt/2; k++ {
				out1 = append(out1, fruit)
			}
		} else if cnt < 0 {
			for k := 0; k > cnt/2; k-- {
				out2 = append(out2, fruit)
			}
		}
	}
	sort.Ints(out1)
	sort.Sort(sort.Reverse(sort.IntSlice(out2)))

	var res int64
	for i := range out1 {
		res += int64(min(min(out1[i], out2[i]), minVal*2))
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
