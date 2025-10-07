package p1488avoidfloodsinthecity

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_avoidFlood(t *testing.T) {
	for _, tc := range []struct {
		rains []int
		want  []int
	}{
		{[]int{1, 2, 0, 0, 1, 2}, []int{-1, -1, 2, 1, -1, -1}},
		{[]int{1, 0, 2, 0, 3, 0, 2, 0, 0, 0, 1, 2, 3}, []int{-1, -1, 2, 1, -1, -1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rains), func(t *testing.T) {
			require.Equal(t, tc.want, avoidFlood(tc.rains))
		})
	}
}

func avoidFlood(rains []int) []int {
	n := len(rains)
	dryDays := make([]int, 0, len(rains))
	seenIdx := make(map[int]int)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	var anyLake int
	for i, x := range rains {
		if x == 0 {
			dryDays = append(dryDays, i)
			continue
		}
		anyLake = x

		// If this lake hasn't been seen (filled), then mark it as seen and continue
		if _, exists := seenIdx[x]; !exists {
			seenIdx[x] = i
			continue
		}

		// If the lake has been filled already, find the first dry day that
		// supersedes the time when the lake was filled
		k := seenIdx[x]
		j := sort.SearchInts(dryDays, k)
		if j == len(dryDays) { // no dry day
			return []int{}
		}
		res[dryDays[j]] = x

		// Remove the dry day
		copy(dryDays[j:], dryDays[j+1:])
		dryDays = dryDays[:len(dryDays)-1]
		delete(seenIdx, x)
		seenIdx[x] = i
	}
	for _, j := range dryDays {
		res[j] = anyLake // any lake
	}
	return res
}
