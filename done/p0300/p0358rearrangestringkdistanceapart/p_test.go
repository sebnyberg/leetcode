package p0358rearrangestringkdistanceapart

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rearrangeString(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want string
	}{
		{"a", 1, "a"},
		{"a", 0, "a"},
		{"aabbcc", 3, "abcabc"},
		{"aaabc", 3, ""},
		{"aaabc", 2, "abaca"},
		{"aaadbbcc", 2, "abacabcd"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, rearrangeString(tc.s, tc.k))
		})
	}
}

func rearrangeString(s string, k int) string {
	// Create a list of character count and positions for which this character
	// is valid
	var count [26]int
	var valid [26]int
	for _, ch := range s {
		count[ch-'a']++
	}

	res := make([]rune, 0, len(s))
	for i := 0; i < len(s); i++ {
		var maxCount int
		index := -1
		for ch := 0; ch < 26; ch++ {
			if valid[ch] <= i && count[ch] > maxCount {
				index = ch
				maxCount = count[ch]
			}
		}
		if index == -1 {
			return ""
		}
		res = append(res, rune(index+'a'))
		valid[index] += k
		count[index]--
	}

	return string(res)
}
