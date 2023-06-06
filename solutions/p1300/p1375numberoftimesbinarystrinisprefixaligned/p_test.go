package p1375numberoftimesbinarystrinisprefixaligned

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numTimesAllBlue(t *testing.T) {
	for i, tc := range []struct {
		flips []int
		want  int
	}{
		{[]int{3, 2, 4, 1, 5}, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numTimesAllBlue(tc.flips))
		})
	}
}

func numTimesAllBlue(flips []int) int {
	// Here's a naive but OK implementation.
	// Keep a boolean array the same size as flips where aligned[i] is true if
	// [0,i] is set to 1.
	m := 1
	n := len(flips)
	for m < n+1 {
		m <<= 1
	}
	tree := make([]int, m*2)
	var query func(i, lo, hi, qlo, qhi int) int
	query = func(i, lo, hi, qlo, qhi int) int {
		if qhi < lo || qlo > hi {
			return 0
		}
		if lo >= qlo && hi <= qhi {
			return tree[i]
		}
		mid := lo + (hi-lo)/2
		return query(i*2, lo, mid, qlo, qhi) + query(i*2+1, mid+1, hi, qlo, qhi)
	}
	update := func(i int) {
		tree[m+i] = 1
		for j := (m + i) / 2; j >= 1; j /= 2 {
			tree[j] = tree[j*2] + tree[j*2+1]
		}
	}
	var res int
	for i := range flips {
		update(flips[i])
		if query(1, 0, m-1, 1, i+1) == i+1 {
			res++
		}
	}
	return res
}
