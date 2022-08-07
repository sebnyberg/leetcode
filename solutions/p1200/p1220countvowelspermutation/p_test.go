package p1220countvowelspermutation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countVowelPermutation(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{1, 5},
		{2, 10},
		{5, 68},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countVowelPermutation(tc.n))
		})
	}
}

func countVowelPermutation(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1
	const mod = 1e9 + 7
	for k := 2; k <= n; k++ {
		aa := (e + i + u) % mod
		ee := (a + i) % mod
		ii := (e + o) % mod
		uu := (i + o) % mod
		oo := i % mod
		a, e, i, o, u = aa, ee, ii, oo, uu
	}
	return (a + e + i + o + u) % mod
}
