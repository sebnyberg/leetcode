package p1234replacethesubstringforbalancedstring

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_balancedString(t *testing.T) {
	for i, tc := range []struct {
		s    string
		want int
	}{
		{"QQWE", 1},
		{"QWER", 0},
		{"QQQW", 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, balancedString(tc.s))
		})
	}
}

func balancedString(s string) int {
	idx := [...]int8{
		'Q': 0,
		'W': 1,
		'E': 2,
		'R': 3,
	}
	var count [4]int
	for _, ch := range s {
		count[idx[ch]]++
	}
	// want[i] = number of occurrences of a character needed within the
	// substring to be able to make the string valid
	var want [4]int
	var ok int
	for x, cnt := range count {
		want[x] = cnt - len(s)/4
		if want[x] <= 0 {
			ok++
		}
	}
	if ok == 4 {
		return 0
	}
	var j int
	res := math.MaxInt32
	for i := range s {
		k := idx[s[i]]
		want[k]--
		if want[k] == 0 {
			ok++
		}
		for j < i && want[idx[s[j]]] < 0 {
			want[idx[s[j]]]++
			j++
		}
		if ok == 4 {
			res = min(res, i-j+1)
			if want[idx[s[j]]] == 0 {
				ok--
			}
			want[idx[s[j]]]++
			j++
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
