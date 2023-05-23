package leetcode_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func TestSegtreeMin(t *testing.T) {
	for i, tc := range []struct {
		arr  []int
		i, j int
		want int
	}{
		{
			[]int{7, 1, -10, 13, 2, 61, 4},
			0, 6, -10,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			var tree leetcode.SegTree[int]
			tree.Init(tc.arr, min, math.MaxInt64)
			got := tree.QueryRange(tc.i, tc.j)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestSegTreeRand(t *testing.T) {
	n := 12938
	ub := 921837
	lb := -128993
	arr := make([]int, n)
	for i := range arr {
		x := rand.Intn(ub - lb + 1)
		arr[i] = x + lb
	}
	var tree leetcode.SegTree[int]
	tree.Init(arr, min, math.MaxInt32)
	for k := 0; k < 4*(n/2); k++ {
		lo := rand.Intn(n)
		hi := rand.Intn(n)
		if lo > hi {
			lo, hi = hi, lo
		}
		want := minRange(arr, lo, hi)
		got := tree.QueryRange(lo, hi)
		if want != got {
			msg := fmt.Sprintf(
				"min query for range %v gave the wrong result\n"+
					"expected: %v\n"+
					"got: %v",
				arr[lo:hi+1],
				want,
				got,
			)
			t.Fatal(msg)
		}
	}
}

func minRange(arr []int, i, j int) int {
	res := arr[i]
	for k := i + 1; k <= j; k++ {
		res = min(res, arr[k])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
