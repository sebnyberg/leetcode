package p1794countpairsofequalsubstringswithmindiff

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countQuadruples(t *testing.T) {
	for _, tc := range []struct {
		firstString  string
		secondString string
		want         int
	}{
		{"abcd", "bccda", 1},
		{"ab", "cd", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.firstString), func(t *testing.T) {
			require.Equal(t, tc.want, countQuadruples(tc.firstString, tc.secondString))
		})
	}
}

func countQuadruples(firstString string, secondString string) int {
	minDist := int(2e5 + 1)
	count := 0
	for c := 'a'; c <= 'z'; c++ {
		j := strings.IndexRune(firstString, c)
		a := strings.LastIndexByte(secondString, byte(c))
		if j == -1 || a == -1 {
			continue
		}
		if j-a < minDist {
			count = 1
			minDist = j - a
		} else if j-a == minDist {
			count++
		}
	}
	return count
}
