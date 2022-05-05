package p1880checkifwordequalssummationoftwowords

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSumEqual(t *testing.T) {
	for _, tc := range []struct {
		firstWord  string
		secondWord string
		targetWord string
		want       bool
	}{
		{"acb", "cba", "cdb", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.firstWord), func(t *testing.T) {
			require.Equal(t, tc.want, isSumEqual(tc.firstWord, tc.secondWord, tc.targetWord))
		})
	}
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	first := ""
	for _, ch := range firstWord {
		first += string(ch - 'a' + '0')
	}
	second := ""
	for _, ch := range secondWord {
		second += string(ch - 'a' + '0')
	}
	target := ""
	for _, ch := range targetWord {
		target += string(ch - 'a' + '0')
	}
	n1, _ := strconv.Atoi(first)
	n2, _ := strconv.Atoi(second)
	n3, _ := strconv.Atoi(target)
	return n1+n2 == n3
}
