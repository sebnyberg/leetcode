package p1625lexicographicallysmallestststringafterapplyingoperations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLexSmallestString(t *testing.T) {
	for _, tc := range []struct {
		s    string
		a    int
		b    int
		want string
	}{
		// {"5525", 9, 2, "2050"},
		{"74", 5, 1, "24"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, findLexSmallestString(tc.s, tc.a, tc.b))
		})
	}
}

func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	bs := []byte(s)
	for i := range bs {
		bs[i] -= '0'
	}

	// find min shift for offset
	findMinShift := func(offset int) int {
		aa := byte(a)
		j := offset
		if j%2 == 0 {
			j = (j + 1) % n
		}
		ch := bs[j]
		var shift int
		for (ch+aa)%10 < ch {
			shift++
			ch = (ch + aa) % 10
		}
		return shift
	}

	valAt := func(i, offset, shift int) byte {
		j := (i + offset) % n
		ch := int(bs[j])
		if j%2 == 1 {
			ch = (ch + shift*a) % 10
		}
		return byte(ch)
	}

	seen := make([]bool, n)
	res := make([]byte, n)
	copy(res, bs)
outer:
	for k := 0; !seen[k]; k = (k + b) % n {
		seen[k] = true
		shift := findMinShift(k)

		for i := range res {
			if x := valAt(i, k, shift); x < res[i] {
				// This solution is optimal, break to capture
				break
			} else if x > res[i] {
				// This is not the solution, continue outer loop
				continue outer
			}
		}

		// Capture result
		for i := range res {
			res[i] = valAt(i, k, shift)
		}
	}

	// Shift back to alpha range
	for i := range res {
		res[i] += '0'
	}

	return string(res)
}
