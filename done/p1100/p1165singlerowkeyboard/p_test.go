package p1165singlerowkeyboard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_calculateTime(t *testing.T) {
	for _, tc := range []struct {
		keyboard string
		word     string
		want     int
	}{
		{"abcdefghijklmnopqrstuvwxyz", "cba", 4},
		{"pqrstuvwxyzabcdefghijklmno", "leetcode", 73},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.keyboard, tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, calculateTime(tc.keyboard, tc.word))
		})
	}
}

func calculateTime(keyboard string, word string) int {
	var indices [26]int
	for i, ch := range keyboard {
		indices[ch-'a'] = i
	}
	var pos int
	var time int
	for _, ch := range word {
		time += abs(indices[ch-'a'] - pos)
		pos = indices[ch-'a']
	}
	return time
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
