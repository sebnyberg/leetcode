package p0828countuniquecharactersofallsubstringsofagivenstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_uniqueLetterString(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"ABA", 8},
		{"ABC", 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, uniqueLetterString(tc.s))
		})
	}
}

func uniqueLetterString(s string) int {
	var charPos [26][2]int
	for i := range charPos {
		charPos[i] = [2]int{-1, -1}
	}
	var res int
	var sum int
	for i := range s {
		c := int(s[i] - 'A')
		// This character adds a unique character to any prior substring which does
		// not include it.
		sum += i - charPos[c][1]

		// And removes a unique character from any substring between the prior and
		// one before that.
		sum -= charPos[c][1] - charPos[c][0]
		charPos[c][0] = charPos[c][1]
		charPos[c][1] = i

		res += sum
	}
	return res
}
