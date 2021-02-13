package p0212wordsearch2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findWords(t *testing.T) {
	for _, tc := range []struct {
		board [][]byte
		words []string
		want  []string
	}{
		// {[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}, []string{"eat", "oath"}},
		// {[][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"abcb"}, []string{"eat", "oath"}},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.board, tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, findWords(tc.board, tc.words))
		})
	}
}

func findWords(board [][]byte, words []string) []string {
	return nil
}
