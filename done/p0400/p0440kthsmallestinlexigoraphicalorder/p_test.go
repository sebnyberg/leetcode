package p0440kthsmallestinlexigoraphicalorder

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLexSortedBelow(t *testing.T) {
	num := 99
	strs := make([]string, 0, num)
	for n := 1; n <= num; n++ {
		strs = append(strs, strconv.Itoa(n))
	}
	sort.Strings(strs)
	x := 22
	res := findLexSortedBelow(x)
	var i int
	for _, s := range strs {
		if s == strconv.Itoa(x) {
			break
		}
		i++
	}
	require.Equal(t, i, res)
}

func Test_findKthNumber(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want int
	}{
		{10, 3, 2},
		{13, 2, 10},
		{1, 1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findKthNumber(tc.n, tc.k))
		})
	}
}

func findKthNumber(n int, k int) int {
	// Case 1: n comes before k in the lexicographically sorted list
	below := findLexSortedBelow(n)
	if k > below+1 {
		max := 1
		for n > 0 {
			max *= 10
			n /= 10
		}
		max = (max / 10) - 1
		return findKthNumber(max-1, k-below-1)
	}
	// Case 2: search in range [1, n]
	fac := 1
	var nbelow int
	for x := n; x > 0; x /= 10 {
		nbelow = 10*nbelow + 10
		fac *= 10
	}
	nbelow = nbelow/10 - 1
	fac /= 10
	res := 1
	for k > 0 {
		if nbelow == 0 {
			res += k - 1
			break
		}
		if k-nbelow < 0 {
			k--
			fac /= 10
			nbelow /= 10
			nbelow--
			res *= 10
		} else {
			k -= nbelow
			res++
		}
	}

	return res
}

func findLexSortedBelow(num int) int {
	var nbelow int
	var res int
	for num > 0 {
		r := num % 10
		if num < 10 {
			r--
		}
		res = res + r + 1 + r*nbelow
		nbelow = nbelow*10 + 10
		num /= 10
	}
	return res - 1
}
