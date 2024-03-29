package p0940distinctsubsequencesii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_distinctSubseqII(t *testing.T) {
	for i, tc := range []struct {
		s    string
		want int
	}{
		{"lee", 5},
		{"abc", 7},
		{"aba", 6},
		{"aaa", 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, distinctSubseqII(tc.s))
		})
	}
}

func distinctSubseqII(s string) int {
	var prevCount [26]int
	for i := range prevCount {
		prevCount[i] = -1
	}
	const mod = 1e9 + 7
	var prev int
	var res int
	for _, ch := range s {
		res = (res*2 - prevCount[ch-'a'] + mod) % mod
		prevCount[ch-'a'] = prev
		prev = res
	}
	return res % mod
}
