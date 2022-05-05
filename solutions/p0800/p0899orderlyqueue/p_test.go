package p0899orderlyqueue

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_orderlyQueue(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want string
	}{
		{"cba", 1, "acb"},
		{"baaca", 3, "aaabc"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, orderlyQueue(tc.s, tc.k))
		})
	}
}

func orderlyQueue(s string, k int) string {
	// If k == 1, then the only available action is to rotate s until reaching
	// the lexicographically smallest value
	if k == 1 {
		bs := make([]byte, len(s)*2)
		copy(bs, s)
		copy(bs[len(s):], s)
		smallest := s
		for i := 1; i < len(s); i++ {
			if cand := string(bs[i : i+len(s)]); cand < smallest {
				smallest = cand
			}
		}
		return smallest
	}
	// k > 0, any combination is possible. Sort string lexicographically
	var charCount [26]int
	for _, ch := range s {
		charCount[ch-'a']++
	}
	res := make([]byte, 0, len(s))
	for ch, count := range charCount {
		res = append(res, bytes.Repeat([]byte{byte(ch) + 'a'}, count)...)
	}
	return string(res)
}
