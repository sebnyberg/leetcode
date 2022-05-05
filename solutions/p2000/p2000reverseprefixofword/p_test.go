package p2000reverseprefixofword

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reversePrefix(t *testing.T) {
	for _, tc := range []struct {
		word string
		ch   byte
		want string
	}{
		{"abcdefd", 'd', "dcbaefd"},
		{"xyxzxe", 'z', "zxyxxe"},
		{"abcd", 'z', "abcd"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, reversePrefix(tc.word, tc.ch))
		})
	}
}

func reversePrefix(word string, ch byte) string {
	idx := -1
	for i := range word {
		if word[i] == ch {
			idx = i
			break
		}
	}
	if idx == -1 {
		return word
	}
	bs := []byte(word)
	for l, r := 0, idx; l < r; l, r = l+1, r-1 {
		bs[l], bs[r] = bs[r], bs[l]
	}
	return string(bs)
}
