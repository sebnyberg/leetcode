package p0815busroutes

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_numBusesToDestination(t *testing.T) {
	for _, tc := range []struct {
		routes         [][]int
		source, target int
		want           int
	}{
		{
			leetcode.ParseMatrix("[[1,2,7],[3,6,7]]"),
			1, 6, 2,
		},
		{
			leetcode.ParseMatrix("[[7,12],[4,5,15],[6],[15,19],[9,12,13]]"),
			15, 12, -1,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.routes), func(t *testing.T) {
			require.Equal(t, tc.want, numBusesToDestination(tc.routes, tc.source, tc.target))
		})
	}
}

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	busesAtStop := make(map[int][]int)
	for i, r := range routes {
		for j := range r {
			busesAtStop[r[j]] = append(busesAtStop[r[j]], i)
		}
	}
	curr := []int{}
	nbuses := len(routes)
	seen := make([]bool, nbuses)
	for _, busIdx := range busesAtStop[source] {
		if seen[busIdx] {
			continue
		}
		seen[busIdx] = true
		curr = append(curr, busIdx)
	}
	next := []int{}
	for changes := 1; len(curr) > 0; changes++ {
		next = next[:0]
		for _, busIdx := range curr {
			for i := 0; i < len(routes[busIdx]); i++ {
				stop := routes[busIdx][i]
				if stop == target {
					return changes
				}
				for _, b := range busesAtStop[stop] {
					if seen[b] {
						continue
					}
					next = append(next, b)
					seen[b] = true
				}
			}
		}
		curr, next = next, curr
	}
	return -1
}
