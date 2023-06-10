package p1415thekthlexicographicalstringofallhappystringsoflengthn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getHappyString(t *testing.T) {
	for i, tc := range []struct {
		n    int
		k    int
		want string
	}{
		{3, 9, "cab"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, getHappyString(tc.n, tc.k))
		})
	}
}

func getHappyString(n int, k int) string {
	// Lazy version. Just incement..
	// First form smallest string
	s := "ababababab"[:n]
	incr := func(s string) string {
		b := []byte(s)
		// Find first index that can be increased
		j := len(b) - 1
		for ; j >= 0; j-- {
			if b[j] == 'c' {
				continue
			}
			if b[j] == 'b' {
				if j == 0 || b[j-1] != 'c' {
					b[j] = 'c'
				} else {
					continue
				}
			}
			if b[j] == 'a' {
				if j == 0 || b[j-1] == 'c' {
					b[j] = 'b'
				} else {
					b[j] = 'c'
				}
			}
			for k := j + 1; k < len(s); k++ {
				if b[k-1] == 'a' {
					b[k] = 'b'
				} else {
					b[k] = 'a'
				}
			}
			return string(b)
		}
		return ""
	}
	for i := 2; i <= k; i++ {
		s = incr(s)
		if s == "" {
			return ""
		}
	}
	return s
}
