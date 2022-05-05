package p0097interleavingstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isInterleave(t *testing.T) {
	for _, tc := range []struct {
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
	} {
		t.Run(fmt.Sprintf("%v/%v -> %v", tc.s1, tc.s2, tc.s3), func(t *testing.T) {
			require.Equal(t, tc.want, isInterleave(tc.s1, tc.s2, tc.s3))
		})
	}
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}

	return helper(make(map[[2]int]bool), 0, 0, []rune(s1), []rune(s2), []rune(s3))
}

func helper(mem map[[2]int]bool, i, j int, s1, s2, s3 []rune) bool {
	if i >= len(s1) && j >= len(s2) {
		return true
	}

	if i < len(s1) && s1[i] == s3[i+j] {
		if _, exists := mem[[2]int{i + 1, j}]; !exists {
			if helper(mem, i+1, j, s1, s2, s3) {
				return true
			} else {
				mem[[2]int{i + 1, j}] = false
			}
		}
	}

	if j < len(s2) && s2[j] == s3[i+j] {
		return helper(mem, i, j+1, s1, s2, s3)
	}

	return false
}
