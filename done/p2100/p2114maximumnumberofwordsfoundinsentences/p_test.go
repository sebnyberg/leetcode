package p2114maximumnumberofwordsfoundinsentences

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mostWordsFound(t *testing.T) {
	for _, tc := range []struct {
		sentences []string
		want      int
	}{
		{[]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}, 6},
		{[]string{"please wait", "continue to fight", "continue to win"}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.sentences), func(t *testing.T) {
			require.Equal(t, tc.want, mostWordsFound(tc.sentences))
		})
	}
}

func mostWordsFound(sentences []string) int {
	var maxLen int
	for _, s := range sentences {
		if n := len(strings.Fields(s)); n > maxLen {
			maxLen = n
		}
	}
	return maxLen
}
