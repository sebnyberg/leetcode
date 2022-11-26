package p1012numberswithrepeateddigits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numDupDigitsAtMostN(t *testing.T) {
	for i, tc := range []struct {
		n    int
		want int
	}{
		{100, 10},
		{1000, 262},
		{20, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numDupDigitsAtMostN(tc.n))
		})
	}
}

func numDupDigitsAtMostN(n int) int {
	// We can turn this problem around and say: how many ways can we form a
	// digit of length m such that there are no repeated digits, and the total
	// value is less than n?

	// First, count numbers which have less digits than n
	x := 9
	facs := []int{9}
	var res int
	for alts := 9; x < n; alts-- {
		x = (x+1)*10 - 1
		res += facs[len(facs)-1]
		facs = append(facs, facs[len(facs)-1]*alts)
	}

	// Then, for each valid prefix, count valid alternatives for the other
	// digits.
	//
	mem := make(map[key]int)
	s := fmt.Sprint(n)
	var bm int
	for i := range s {
		for x := 0; x < int(s[i]-'0'); x++ {
			if i == 0 && x == 0 {
				continue
			}
			if bm&(1<<x) > 0 {
				continue
			}
			a := dp(mem, bm|(1<<x), len(s)-i-1)
			res += a
		}
		if bm&(1<<int(s[i]-'0')) > 0 {
			break
		}
		bm |= (1 << int(s[i]-'0'))
	}
	// Check n as well
	ok := func(ss string) bool {
		var bm int
		for i := range ss {
			if bm&(1<<int(ss[i]-'0')) > 0 {
				return false
			}
			bm |= (1 << int(ss[i]-'0'))
		}
		return true
	}
	if ok(s) {
		res++
	}

	actual := n - res
	return actual
}

func dp(mem map[key]int, bm, ndigits int) int {
	k := key{
		bm:     bm,
		digits: ndigits,
	}
	if v, exists := mem[k]; exists {
		return v
	}
	if ndigits == 0 {
		return 1
	}
	var res int
	for x := 0; x <= 9; x++ {
		if bm&(1<<x) > 0 {
			continue
		}
		res += dp(mem, bm|(1<<x), ndigits-1)
	}
	mem[k] = res
	return res
}

type key struct {
	bm     int
	digits int
}
