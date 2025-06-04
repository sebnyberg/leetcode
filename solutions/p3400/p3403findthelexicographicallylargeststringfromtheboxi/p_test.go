package p3403findthelexicographicallylargeststringfromtheboxi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_answerString(t *testing.T) {
	for _, tc := range []struct {
		word       string
		numFriends int
		want       string
	}{
		{"gh", 1, "gh"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, answerString(tc.word, tc.numFriends))
		})
	}
}

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	maxLen := len(word) - numFriends + 1
	res := word[:1]
	for i := range word {
		s := word[i:min(i+maxLen, len(word))]
		m := min(len(s), len(res))
		if s[:m] < res[:m] {
			continue
		}
		if s[:m] > res[:m] {
			res = s
			continue
		}
		// s[:m] == res[:m]
		if len(s) > len(res) {
			res = s
		}
	}
	return res
}
