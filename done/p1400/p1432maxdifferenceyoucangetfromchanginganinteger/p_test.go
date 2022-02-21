package p1432maxdifferenceyoucangetfromchanginganinteger

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxDiff(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want int
	}{
		{123456, 820000},
		{555, 888},
		{9, 8},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, maxDiff(tc.num))
		})
	}
}

func maxDiff(num int) int {
	parse := func(s string) int {
		var res int
		for i := range s {
			res *= 10
			res += int(s[i] - '0')
		}
		return res
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	s := fmt.Sprint(num)
	var maxDiff int
	for _, x := range "0123456789" {
		a := strings.ReplaceAll(s, string(x), "9")
		for _, y := range "0123456789" {
			ch := "0"
			if rune(s[0]) == y {
				ch = "1"
			}
			b := strings.ReplaceAll(s, string(y), ch)
			maxDiff = max(maxDiff, parse(a)-parse(b))
		}
	}

	return maxDiff
}
