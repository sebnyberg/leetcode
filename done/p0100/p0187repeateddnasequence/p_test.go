package p0187repeateddnasequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want []string
	}{
		{"AAAAAAAAAAA", []string{"AAAAAAAAAA"}},
		{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
		{"AAAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, findRepeatedDnaSequences(tc.s))
		})
	}
}

func findRepeatedDnaSequences(s string) []string {
	seen := make(map[string]byte)
	res := make([]string, 0)
	for i := 0; i < len(s)-9; i++ {
		k := s[i : i+10]
		if seen[k] == 2 {
			continue
		}
		if seen[k] == 1 {
			res = append(res, k)
		}
		seen[k]++
	}

	return res
}
