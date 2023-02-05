package p2555mazimizewinfromtwosegments

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximizeWin(t *testing.T) {
	for i, tc := range []struct {
		prizePositions []int
		k              int
		want           int
	}{
		{[]int{1, 1, 2, 2, 3, 3, 5}, 2, 7},
		{[]int{1, 2, 3, 4}, 0, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maximizeWin(tc.prizePositions, tc.k))
		})
	}
}

func maximizeWin(prizePositions []int, k int) int {
	// Do a left-right combination scan, where each position reports the maximum
	// total number of prizes to the left/right of that position.
	//
	// First, we must coalesce the positions into a single slice with counts and
	// positions
	var count []int
	var pos []int
	sort.Ints(prizePositions)
	count = append(count, 1)
	pos = append(pos, prizePositions[0])
	var j int
	for _, x := range prizePositions[1:] {
		if x != pos[j] {
			pos = append(pos, x)
			count = append(count, 0)
			j++
		}
		count[j]++
	}
	// fmt.Println("count", count)
	// fmt.Println("pos", pos)

	// left[i] maximum number of prizes ending before pos[i]
	var l int
	var res int
	n := len(pos)
	left := make([]int, n+1)
	maxLeft := make([]int, n+1)
	for i := range count {
		res += count[i]
		for pos[i]-pos[l] > k {
			res -= count[l]
			l++
		}
		left[i+1] = res
		maxLeft[i+1] = max(maxLeft[i], res)
	}
	// fmt.Println("left", left)
	// fmt.Println("maxLeft", maxLeft)

	right := make([]int, n)
	maxRight := make([]int, n+1)
	res = 0
	r := n - 1
	for i := n - 1; i >= 0; i-- {
		res += count[i]
		for pos[r]-pos[i] > k {
			res -= count[r]
			r--
		}
		right[i] = res
		maxRight[i] = max(maxRight[i+1], res)
	}
	// fmt.Println("right", right)
	// fmt.Println("maxRight", maxRight)
	var ret int
	for i := range maxLeft {
		ret = max(ret, maxLeft[i]+maxRight[i])
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
