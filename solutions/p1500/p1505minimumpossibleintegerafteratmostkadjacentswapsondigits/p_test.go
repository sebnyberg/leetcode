package p1505minimumpossibleintegerafteratmostkadjacentswapsondigits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minInteger(t *testing.T) {
	for i, tc := range []struct {
		num  string
		k    int
		want string
	}{
		{"4321", 4, "1342"},
		{"100", 1, "010"},
		{"36789", 1000, "36789"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minInteger(tc.num, tc.k))
		})
	}
}

func minInteger(num string, k int) string {
	// If a series of swaps can reduce the leftmost digit, then a series that
	// does not can be discarded.
	//
	// Other than that, there exist some degree of creative freedom. The
	// question is: does the order of swaps performed before moving the optimal
	// digit to the leftmost position?
	//
	// Let's say we have num=254321, and k=10. Would we ever want to do anything
	// but shuffle the number to 15432? The answer is: no, any move we could
	// make before shuffling the small number in front could just as well be
	// done later with no change in outcome.
	//
	// This tells us that the problem lies in efficiently finding the smallest
	// value within the range [i+1, min(i+k, n-1)] and shuffling it to the
	// front.
	//
	// Since there are only 9 digits, we can create a list of digits and their
	// original position. Once we start taking digits however, there will be an
	// issue with keeping track of the actual distance given that positions
	// change as numbers are swapped to the front.
	//
	// This is where a segment tree is useful. With a segment tree marking 1s
	// where numbers are still in place, and 0 where they are not, we can update
	// the relative distances to numbers in the number in O(logn).
	//
	n := 1
	for n < len(num) {
		n *= 2
	}

	// Collect lists of digits' relative positions
	var idx [10][]int
	for i, x := range num {
		idx[x-'0'] = append(idx[x-'0'], i)
	}

	// Initialize segment tree
	segtree := make([]int, 2*n)
	for i := range num {
		segtree[n+i] = 1
	}
	for i := n - 1; i >= 1; i-- {
		segtree[i] = segtree[2*i] + segtree[2*i+1]
	}

	res := make([]byte, len(num))
	for i := range num {
		for x := 0; x <= 9; x++ {
			if len(idx[x]) == 0 {
				continue
			}
			var dist int
			hi := n - 1
			m := idx[x][0]
			for j := 1; j < n; {
				if m <= hi/2 {
					j = j * 2
				} else {
					dist += segtree[j*2]
					m -= (hi / 2) + 1
					j = j*2 + 1
				}
				hi /= 2
			}
			// note: the final node in the tree wasn't counted, that's why the
			// 'distance' - it's the index delta
			if dist > k {
				continue
			}
			k -= dist

			// update segtree
			for j := n + idx[x][0]; j >= 1; j /= 2 {
				segtree[j]--
			}

			idx[x] = idx[x][1:]
			res[i] = byte(x + '0')
			break
		}
	}
	return string(res)
}

func min(a, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}
