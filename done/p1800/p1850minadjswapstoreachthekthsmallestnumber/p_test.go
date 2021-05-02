package p1850minadjswapstoreachthekthsmallestnumber

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getMinSwaps(t *testing.T) {
	for _, tc := range []struct {
		num  string
		k    int
		want int
	}{
		{"11112", 4, 4},
		{"5489355142", 4, 2},
		{"00123", 1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, getMinSwaps(tc.num, tc.k))
		})
	}
}

func getMinSwaps(num string, k int) int {
	// 1. Find kth smallest number greater than num
	numCpy := num
	n := len(numCpy)
	for i := 0; i < k; i++ {
		numCpy = findNext(numCpy, n)
	}

	// 2. Calculate number of swaps required to reach number
	orig := []byte(num)
	target := []byte(numCpy)
	var nswaps int
	for i := range orig {
		if orig[i] != target[i] {
			var swaps int
			for swaps = 1; orig[i+swaps] != target[i]; swaps++ {
			}
			for j := swaps - 1; j >= 0; j-- {
				orig[i+j], orig[i+j+1] = orig[i+j+1], orig[i+j]
				nswaps++
			}
		}
	}
	return nswaps
}

func findNext(num string, n int) string {
	bs := []byte(num)
	var i int
	for i = n - 2; i >= 0; i-- {
		if bs[i] < bs[i+1] {
			break
		}
	}
	// find smallest value greater than the one to swap
	firstGreaterIdx := i + 1
	firstGreater := bs[i+1]
	for j := i + 1; j < n; j++ {
		if bs[j] > bs[i] && bs[j] < firstGreater {
			firstGreater = bs[j]
			firstGreaterIdx = j
		}
	}
	bs[i], bs[firstGreaterIdx] = bs[firstGreaterIdx], bs[i]
	sort.Slice(bs[i+1:], func(j, k int) bool {
		return bs[i+1+j] < bs[i+1+k]
	})
	return string(bs)
}
