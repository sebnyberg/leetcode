package p2047numberofvalidwordsinasentence

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countValidWords(t *testing.T) {
	for _, tc := range []struct {
		sentence string
		want     int
	}{
		{"cat and dog", 3},
		{"!this  1-s b8d!", 0},
		{"alice and  bob are playing stone-game10", 5},
		{"a-!", 0},
		{"he bought 2 pencils, 3 erasers, and 1  pencil-sharpener.", 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.sentence), func(t *testing.T) {
			require.Equal(t, tc.want, countValidWords(tc.sentence))
		})
	}
}

func countValidWords(sentence string) int {
	fields := strings.Fields(sentence)
	var validCount int
	for _, word := range fields {
		var hyphenCount int
		for i, ch := range word {
			switch {
			case ch == '-':
				if i == 0 || i == len(word)-1 || hyphenCount == 1 {
					goto continueLoop
				}
				hyphenCount++
			case ch >= 'a' && ch <= 'z':
			case ch >= '0' && ch <= '9':
				goto continueLoop
			case ch == '!' || ch == '.' || ch == ',':
				if i != len(word)-1 || len(word) > 1 && word[len(word)-2] == '-' {
					goto continueLoop
				}
			}
		}
		validCount++
	continueLoop:
	}
	return validCount
}
