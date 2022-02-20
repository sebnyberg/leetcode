package p1100findklengthsubstringswithnorepeatedcharacters

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numKLenSubstrNoRepeats(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want int
	}{
		{},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, numKLenSubstrNoRepeats(tc.s, tc.k))
		})
	}
}

func numKLenSubstrNoRepeats(s string, k int) int {
	var res int
	var count [26]int
	var nonUnique int
	for i, ch := range s {
		count[ch-'a']++
		if count[ch-'a'] == 2 {
			nonUnique++
		}
		if i >= k {
			idx := int(s[i-k] - 'a')
			count[idx]--
			if count[idx] == 1 {
				nonUnique--
			}
		}
		if i >= k-1 && nonUnique == 0 {
			res++
		}
	}
	return res
}
