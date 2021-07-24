package p0187repeateddnaseqrollinghash

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
		{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
		{"AAAAAAAAAAA", []string{"AAAAAAAAAA"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, findRepeatedDnaSequences(tc.s))
		})
	}
}

func findRepeatedDnaSequences(s string) []string {
	hashval := map[byte]int{
		'A': 0,
		'C': 1,
		'G': 2,
		'T': 3,
	}
	finalLetter := 1 << 18
	h := 0
	seen := make(map[int]bool)
	res := make([]string, 0)
	for i := range s {
		if i >= 10 {
			h -= hashval[s[i-10]] * finalLetter
		}
		// Add new letter
		h *= 4
		h += hashval[s[i]]
		if i >= 9 {
			if v, exists := seen[h]; !exists {
				seen[h] = true
			} else if v {
				res = append(res, s[i-9:i+1])
				seen[h] = false
			}
		}
	}

	return res
}
