package p1461checkifastringcontainsallbinarycodesofsizek

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hasAllCodes(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want bool
	}{
		{"001010101010110", 4, false},
		{"00110110", 2, true},
		{"00110", 2, true},
		{"0110", 1, true},
		{"0110", 2, false},
		{"0000000001011100", 4, false},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.s, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, hasAllCodes(tc.s, tc.k))
		})
	}
}

func hasAllCodes(s string, k int) bool {
	h1 := 0
	pow := 1 << (k - 1)
	// n := len(s)
	seen := make(map[int]struct{})
	for i, ch := range s {
		h1 = (h1*2 + int(ch-'0'))
		if i < k-1 {
			continue
		}
		seen[h1] = struct{}{}
		// Remove "oldest" element
		h1 = (h1 - int(s[i-k+1]-'0')*pow)
	}
	for n := 0; n < 1<<k; n++ {
		if _, exists := seen[n]; !exists {
			return false
		}
	}
	return true
}
