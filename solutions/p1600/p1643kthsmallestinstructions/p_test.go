package p1643kthsmallestinstructions

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthSmallestPath(t *testing.T) {
	for i, tc := range []struct {
		destionation []int
		k            int
		want         string
	}{
		{[]int{2, 3}, 2, "HHVHV"},
		{[]int{1, 3}, 3, "HVHH"},
		{[]int{2, 2}, 6, "VVHH"},
		{[]int{15, 15}, 155117520, "VVVVVVVVVVVVVVVHHHHHHHHHHHHHHH"},
		{[]int{2, 3}, 3, "HHVVH"},
		{[]int{2, 3}, 1, "HHHVV"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, kthSmallestPath(tc.destionation, tc.k))
		})
	}
}

func kthSmallestPath(destination []int, k int) string {
	b := newBin(destination[1] + destination[0])
	res := f(b, destination[1], destination[0], k)
	res = strings.ReplaceAll(res, "A", "H")
	res = strings.ReplaceAll(res, "B", "V")
	return res
}

func f(b bin, na, nb, k int) string {
	// Any valid solution takes w-1 steps right and h-1 steps down.
	// Let's use A and B instead- much easier to compare.
	if k == 1 {
		var res string
		for i := 0; i < na; i++ {
			res += "A"
		}
		for i := 0; i < nb; i++ {
			res += "B"
		}
		return res
	}
	k--

	// We start with AAAAA...BBBB
	//
	// Then by moving the leftmost  B to the left, we open up for (nb+1 over 1)
	// permutations, i.e. AAAA...B(BBBBA) and AAAA...B(BBBAB) etc.
	//
	// Then by moving the B one step further, we have (bs+1 over 2) permutations.
	// Once the total number of permutations for a move would exceed the
	// requested k, then we can reduce k by that amount and do a recursive call.
	for as := 1; as <= na; as++ {
		ways := b.binomial(nb-1+as, as)
		if ways == k {
			var res string
			for i := 0; i < na-as; i++ {
				res += "A"
			}
			for i := 0; i < nb; i++ {
				res += "B"
			}
			for i := 0; i < as; i++ {
				res += "A"
			}
			return res
		}
		if ways > k {
			// The correct configuration starts with
			// (na - as) A's, then one B, then some set of As and Bs.
			// We find the correct one via a recursive call
			var res string
			for i := 0; i < (na - as); i++ {
				res += "A"
			}
			res += "B"
			return res + f(b, as, nb-1, k)
		}
		k -= ways
	}
	panic("something is srsly wrong")
}

type bin struct {
	vals [][]int
}

func newBin(n int) bin {
	var b bin
	b.vals = make([][]int, n+1)
	for i := range b.vals {
		b.vals[i] = make([]int, n+1)
	}
	for i := range b.vals {
		b.vals[i][0] = 1
		b.vals[i][i] = 1
	}
	for m := 1; m <= n; m++ {
		for k := 1; k <= m-1; k++ {
			b.vals[m][k] = b.vals[m-1][k-1] + b.vals[m-1][k]
		}
	}
	return b
}

func (b bin) binomial(n, k int) int {
	return b.vals[n][k]
}
