package p2940findbuildingwherealiceandbobcanmeet

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_leftmostBuildingQueries(t *testing.T) {
	for i, tc := range []struct {
		heights []int
		queries [][]int
		want    []int
	}{
		{[]int{6, 4, 8, 5, 2, 7}, [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}}, []int{2, 5, -1, 5, 2}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, leftmostBuildingQueries(tc.heights, tc.queries))
		})
	}
}

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	// Each query can be reformulated as...
	//
	// Given a query [i, j], what is the smallest height[k] such that
	// either k > max(i, j) and height[k] > max(height[i], height[j])
	// or k == max(i, j) and height[max(i,j)] > height[min(i, j)]
	//
	// There is some trick to use here where we consider either houses in order
	// or queries in order of position.
	//
	// For example, we could consider houses and queries by descending (max)
	// height. Adding houses to a segment tree and allowing a "find leftmost"
	// query would allow us to quickly find the leftmost index of a height that
	// is higher than or equal to the highest height of the query and has an
	// index larger than the rightmost indexed person.
	//
	m := len(heights)
	n := 1
	for n < m {
		n <<= 2
	}
	t := make([]int, n*2)
	for i := range t {
		t[i] = math.MaxInt32
	}

	var query func(i, tl, tr, l, r int) int
	query = func(i, tl, tr, l, r int) int {
		if r < tl || l > tr {
			return math.MaxInt32
		}
		if tl >= l && tr <= r {
			return t[i]
		}
		mid := tl + (tr-tl)/2
		a := query(i*2, tl, mid, l, r)
		b := query(i*2+1, mid+1, tr, l, r)
		return min(a, b)
	}

	// Create index slice sorted by maximum height of player positions in the
	// query
	idx := make([]int, len(queries))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		a := queries[idx[i]]
		b := queries[idx[j]]
		return max(heights[a[0]], heights[a[1]]) > max(heights[b[0]], heights[b[1]])
	})

	heightIdx := make([]int, len(heights))
	for i := range heightIdx {
		heightIdx[i] = i
	}
	sort.Slice(heightIdx, func(i, j int) bool {
		return heights[heightIdx[i]] > heights[heightIdx[j]]
	})

	var k int
	res := make([]int, len(queries))
	for _, j := range idx {
		q := queries[j]
		if q[0] == q[1] {
			res[j] = q[0]
			continue
		}
		if q[0] < q[1] && heights[q[0]] < heights[q[1]] {
			res[j] = q[1]
			continue
		}
		if q[1] < q[0] && heights[q[1]] < heights[q[0]] {
			res[j] = q[0]
			continue
		}
		hh := max(heights[q[0]], heights[q[1]])
		for k < len(heights) && heights[heightIdx[k]] > hh {
			// Add to segment tree
			t[n+heightIdx[k]] = heightIdx[k]
			for x := (n + heightIdx[k]) / 2; x >= 1; x /= 2 {
				t[x] = min(t[x*2], t[x*2+1])
			}
			k++
		}

		// Query for lowest indexed building such that the index >= i
		pos := query(1, 0, n-1, max(q[0], q[1]), n-1)
		if pos == math.MaxInt32 {
			res[j] = -1
		} else {
			res[j] = pos
		}
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
