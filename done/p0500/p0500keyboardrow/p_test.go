package p0500keyboardrow

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findWords(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  []string
	}{
		{[]string{"Hello", "Alaska", "Dad", "Peace"}, []string{"Alaska", "Dad"}},
		{[]string{"omk"}, nil},
		{[]string{"adsdf", "sfd"}, []string{"adsdf", "sfd"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, findWords(tc.words))
		})
	}
}

func findWords(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	oneRow := func(word string) bool {
		rowIdx := -1
		for _, ch := range word {
			for i, row := range rows {
				if !strings.ContainsRune(row, ch) {
					continue
				}
				if rowIdx == -1 {
					rowIdx = i
				} else if rowIdx != i {
					return false
				}
			}
		}
		return true
	}

	var res []string
	for _, w := range words {
		if oneRow(w) {
			res = append(res, w)
		}
	}
	return res
}
