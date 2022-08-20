package p2381shiftingletters2

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_shiftingLetters(t *testing.T) {
	for _, tc := range []struct {
		s      string
		shifts [][]int
		want   string
	}{
		{"abc", leetcode.ParseMatrix("[[0,1,0],[1,2,1],[0,2,1]]"), "ace"},
		{"dztz", leetcode.ParseMatrix("[[0,0,0],[1,1,1]]"), "catz"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, shiftingLetters(tc.s, tc.shifts))
		})
	}
}

func shiftingLetters(s string, shifts [][]int) string {
	// Shifting back is the same as adding 25
	// Shifting forward is the same as adding 1
	n := len(s)
	deltas := make([]byte, n+1)
	for _, s := range shifts {
		if s[2] == 1 {
			deltas[s[0]] += 1
			deltas[s[1]+1] += 25
		} else {
			deltas[s[0]] += 25
			deltas[s[1]+1] += 1
		}
	}
	var shift byte
	res := make([]byte, n)
	for i := range res {
		shift = (shift + deltas[i]) % 26
		res[i] = byte(((byte(s[i]-'a') + shift) % 26) + 'a')
	}
	return string(res)
}
