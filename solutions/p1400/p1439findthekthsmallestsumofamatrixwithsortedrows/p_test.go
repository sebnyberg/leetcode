package p1439findthekthsmallestsumofamatrixwithsortedrows

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_kthSmallest(t *testing.T) {
	things := make([]string, 40)
	for i := range things {
		things[i] = "["
		for j := 0; j < 39; j++ {
			things[i] += "1,"
		}
		things[i] += "1]"
	}
	thing := strings.Join(things, ",")
	for i, tc := range []struct {
		mat  [][]int
		k    int
		want int
	}{
		{
			leetcode.ParseMatrix("[" + thing + "]"),
			5, 40,
		},
		{
			leetcode.ParseMatrix("[[1,3,11],[2,4,6]]"), 5, 7,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, kthSmallest(tc.mat, tc.k))
		})
	}
}

func kthSmallest(mat [][]int, k int) int {
	// Let's do some calculations.
	// There are up to 40 options per row, then another 40 rows, so 40^40
	// different arrays.
	// Note that since k <= 200, we could add the first element of each row to
	// find the smallest value, then add the smallest second element from a row
	// to find the second smallest, and so on. It may feel inefficient, but
	// simply iterating over each row each time isnt too bad O(k*m)
	curr := map[int]int{}
	next := map[int]int{}
	buf := []int{}
	curr[0] = 1
	for i := range mat {
		for k := range next {
			delete(next, k)
		}
		for _, v := range mat[i] {
			for x, cnt := range curr {
				next[x+v] = min(200, next[x+v]+cnt)
			}
		}
		// Prune sums that are not applicable
		buf = buf[:0]
		for v := range next {
			buf = append(buf, v)
		}
		sort.Ints(buf)
		if len(buf) > k {
			kth := buf[k]
			for v := range next {
				if v >= kth {
					delete(next, v)
				}
			}
		}
		curr, next = next, curr
	}
	type valCount struct {
		val   int
		count int
	}
	var sums []valCount
	for val, cnt := range curr {
		sums = append(sums, valCount{val, cnt})
	}
	sort.Slice(sums, func(i, j int) bool {
		return sums[i].val < sums[j].val
	})
	var count int
	for _, vc := range sums {
		if count+vc.count >= k {
			return vc.val
		}
		count += vc.count
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
