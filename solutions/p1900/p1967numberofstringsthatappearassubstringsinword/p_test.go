package p1967numberofstringsthatappearassubstringsinword

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numOfStrings(t *testing.T) {
	for _, tc := range []struct {
		patterns []string
		word     string
		want     int
	}{
		{[]string{"a", "abc", "bc", "d"}, "abc", 3},
		{[]string{"a", "b", "c"}, "aaaaabbbbb", 2},
		{[]string{"a", "a", "a"}, "ab", 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.patterns), func(t *testing.T) {
			require.Equal(t, tc.want, numOfStrings(tc.patterns, tc.word))
		})
	}
}

func numOfStrings(patterns []string, word string) int {
	var count int
	for _, pat := range patterns {
		for i := 0; i < len(word)-len(pat)+1; i++ {
			if word[i:i+len(pat)] == pat {
				count++
				break
			}
		}
	}
	return count
}
