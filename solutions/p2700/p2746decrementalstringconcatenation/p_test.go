package p2746decrementalstringconcatenation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimizeConcatenatedLength(t *testing.T) {
	for i, tc := range []struct {
		words []string
		want  int
	}{
		{[]string{"aa", "ab", "bc"}, 4},
		{[]string{"ab", "b"}, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimizeConcatenatedLength(tc.words))
		})
	}
}

func minimizeConcatenatedLength(words []string) int {
	// The big "insight" is that we only care about the first and last letter of
	// the current word. So we can memoize the result as a 1000*27*27*1000 array,
	// where dp[i][l][r] = the minimum number of extra characters from position
	// i and onwards given that the previous string starts with l and ends with r
	n := len(words)
	mem := make([][][]int, n)
	for i := range mem {
		mem[i] = make([][]int, 27)
		for j := range mem[i] {
			mem[i][j] = make([]int, 27)
			for k := range mem[i][j] {
				mem[i][j][k] = -1
			}
		}
	}
	l := words[0][0] - 'a'
	r := words[0][len(words[0])-1] - 'a'
	res := len(words[0]) + dp(mem, words, 1, l, r, n)
	return res
}

func dp(mem [][][]int, words []string, i int, l, r byte, n int) int {
	if i == n {
		return 0
	}
	if mem[i][l][r] != -1 {
		return mem[i][l][r]
	}

	ll := words[i][0] - 'a'
	rr := words[i][len(words[i])-1] - 'a'

	// Option 1, if r == words[i][0] then we can strip one character
	opt1 := len(words[i])
	if r == ll {
		opt1--
	}
	res := opt1 + dp(mem, words, i+1, l, rr, n)

	// Option 2, if rr == l, we can strip one character
	opt2 := len(words[i])
	if rr == l {
		opt2--
	}
	res = min(res, opt2+dp(mem, words, i+1, ll, r, n))

	mem[i][l][r] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
