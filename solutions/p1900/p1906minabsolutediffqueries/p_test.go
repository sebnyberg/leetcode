package p1906minabsolutediffqueries

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDifference(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		queries [][]int
		want    []int
	}{
		{[]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}}, []int{2, 1, 4, 1}},
		{[]int{4, 5, 2, 2, 7, 10}, [][]int{{2, 3}, {0, 2}, {0, 5}, {3, 5}}, []int{-1, 1, 1, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minDifference(tc.nums, tc.queries))
		})
	}
}

func minDifference(nums []int, queries [][]int) []int {
	// Create 100 BITs which contain the frequencies for each number.
	trees := make([]tree, 101)
	n := len(nums)
	for i := range trees {
		trees[i] = make(tree, n+1)
	}

	// Add frequency counters
	for i, num := range nums {
		for j := i; j < n; j |= j + 1 {
			trees[num][j] += 1
		}
	}

	// For each query, get the frequency counts for each number and find
	// the shortest distance in the list of numbers
	res := make([]int, len(queries))
	for i, query := range queries {
		l, r := uint32(query[0]), uint32(query[1])
		minDist := math.MaxInt32
		prev := -1000
		for num := 1; num <= 100; num++ {
			freq := trees[num].sumRange(l, r+1)
			if freq > 0 {
				minDist = min(minDist, num-prev)
				prev = num
			}
		}
		if minDist > 500 {
			res[i] = -1
		} else {
			res[i] = minDist
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type tree []uint32

func (t tree) sum(pos uint32) uint32 {
	var sum uint32
	for pos > 0 {
		sum += t[pos-1]
		pos -= pos & -pos
	}
	return sum
}

// sumRange returns the sum from [start,end) (similar to how slice indexing works)
func (t tree) sumRange(start, end uint32) uint32 {
	return t.sum(end) - t.sum(start)
}
