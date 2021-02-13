package p0060permseqv2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getPermutation(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want string
	}{
		{3, 4, "231"},
		{3, 2, "132"},
		{3, 3, "213"},
		{4, 9, "2314"},
		{3, 1, "123"},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.n, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, getPermutation(tc.n, tc.k))
		})
	}
}

func getPermutation(n int, k int) string {
	nperms := 1
	for i := 2; i <= n; i++ {
		nperms *= i
	}
	var taken [9]bool
	res := make([]uint8, 0, n)
	findPerms(&res, taken, nperms, n, k-1)
	for i := range res {
		res[i] += '0' + 1
	}
	return string(res)
}

func findPerms(res *[]uint8, taken [9]bool, maxperms int, n int, k int) {
	if n == 0 {
		return
	}
	maxperms /= n
	// Steps to take from the smallest available number at this position
	steps := k / maxperms
	j := 0
	for {
		// Find smallest available number
		for ; taken[j]; j++ {
		}
		if steps == 0 {
			taken[j] = true
			*res = append(*res, uint8(j))
			findPerms(res, taken, maxperms, n-1, k%maxperms)
			return
		}
		j++
		steps--
	}
}
