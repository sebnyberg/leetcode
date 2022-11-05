package p0903validpermutationsfordisequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numPermsDISequence(t *testing.T) {
	for i, tc := range []struct {
		s    string
		want int
	}{
		{"DID", 5},
		{"D", 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numPermsDISequence(tc.s))
		})
	}
}

const mod = 1e9 + 7

func numPermsDISequence(s string) int {
	var res int
	n := len(s)
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	for i := 0; i <= n; i++ {
		res = (res + numPerms(mem, s, i)) % mod
	}
	return res
}

func numPerms(mem [][]int, s string, x int) int {
	n := len(s)
	if mem[n][x] != -1 {
		return mem[n][x]
	}
	if len(s) == 0 {
		return 1
	}
	var res int
	if s[0] == 'I' {
		for y := x + 1; y <= n; y++ {
			res = (res + numPerms(mem, s[1:], y-1)) % mod
		}
	} else {
		for y := x - 1; y >= 0; y-- {
			res = (res + numPerms(mem, s[1:], y)) % mod
		}
	}
	mem[n][x] = res
	return res
}
