package p2060checkifanoriginalstringexiststgiventwoencodedstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_possiblyEquals(t *testing.T) {
	for _, tc := range []struct {
		s1   string
		s2   string
		want bool
	}{
		{"internationalization", "i18n", true},
		{"l123e", "44", true},
		{"a5b", "c5b", false},
		{"112s", "g841", true},
		{"ab", "a2", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s1), func(t *testing.T) {
			require.Equal(t, tc.want, possiblyEquals(tc.s1, tc.s2))
		})
	}
}

func possiblyEquals(s1 string, s2 string) bool {
	var seen [50][50][2000]bool

	n1, n2 := len(s1), len(s2)
	var visit func(i1, i2, diff int) bool
	visit = func(i1, i2, diff int) bool {
		if i1 == n1 && i2 == n2 {
			return diff == 0
		}
		if seen[i1][i2][diff+1000] {
			return false
		}
		seen[i1][i2][diff+1000] = true

		switch {
		case diff >= 0 && i1 < n1 && s1[i1] <= '9':
			var x int
			for j := i1; j < min(i1+3, n1) && s1[j] <= '9'; j++ {
				x = x*10 + int(s1[j]-'0')
				if visit(j+1, i2, diff-x) {
					return true
				}
			}
		case diff <= 0 && i2 < n2 && s2[i2] <= '9':
			var x int
			for j := i2; j < min(i2+3, n2) && s2[j] <= '9'; j++ {
				x = x*10 + int(s2[j]-'0')
				if visit(i1, j+1, diff+x) {
					return true
				}
			}
		case diff == 0:
			if i1 >= n1 || i2 >= n2 || s1[i1] != s2[i2] {
				return false
			}
			return visit(i1+1, i2+1, 0)
		case diff > 0:
			if i1 >= n1 {
				return false
			}
			return visit(i1+1, i2, diff-1)
		default:
			if i2 >= n2 {
				return false
			}
			return visit(i1, i2+1, diff+1)
		}
		return false
	}

	return visit(0, 0, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
