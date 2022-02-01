package p0424longestrepeatingcharacterreplacement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_characterReplacement(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, characterReplacement(tc.s, tc.k))
		})
	}
}

func characterReplacement(s string, k int) int {
	var count [26]int

	var res, maxCount int
	for i := range s {
		count[s[i]-'A']++
		if c := count[s[i]-'A']; c > maxCount {
			maxCount = c
		}
		if res < maxCount+k {
			res++
		} else {
			count[s[i-res]-'A']--
		}
	}

	return res
}
