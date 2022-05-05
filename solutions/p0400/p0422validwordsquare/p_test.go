package p0422validwordsquare

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validWordSquare(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  bool
	}{
		{[]string{"ball", "asee", "let", "lep"}, false},
		{[]string{"abcd", "bnrt", "crmy", "dtye"}, true},
		{[]string{"abcd", "bnrt", "crm", "dt"}, true},
		{[]string{"ball", "area", "read", "lady"}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, validWordSquare(tc.words))
		})
	}
}

func validWordSquare(words []string) bool {
	var n int
	for _, v := range words {
		n = max(n, len(v))
	}
	if len(words) != n {
		return false
	}
	// Pad words
	for i := range words {
		if len(words[i]) != n {
			words[i] = words[i] + strings.Repeat("*", n-len(words[i]))
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if words[i][j] != words[j][i] {
				return false
			}
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
