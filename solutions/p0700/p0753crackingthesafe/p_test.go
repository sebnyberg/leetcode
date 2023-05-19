package p0753crackingthesafe

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_crackSafe(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want string
	}{
		{1, 2, "01"},
		{2, 2, "00110"},
		{6, 2, "000000100001100010100011100100101100110100111101010111011011111100000"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, crackSafe(tc.n, tc.k))
		})
	}
}

const alphabet = "0123456789"

func crackSafe(n int, k int) string {
	lyndonWords := make([]string, 0)
	safeAlphabet := alphabet[:k]
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			lyndonWords = append(lyndonWords, duval(safeAlphabet, i)...)
		}
	}
	sort.Strings(lyndonWords)
	return strings.Join(lyndonWords, "") + strings.Repeat("0", n-1)
}

// https://www.geeksforgeeks.org/generating-lyndon-words-of-length-n/
func duval(alphabet string, maxLen int) []string {
	alphaBytes := []byte(alphabet)
	sort.Slice(alphaBytes, func(i, j int) bool {
		return alphaBytes[i] < alphaBytes[j]
	})
	res := make([]string, 0)
	wordIndices := []int{-1}
	alphaLen := len(alphabet)
	n := len(wordIndices)
	for len(wordIndices) > 0 {
		wordIndices[n-1] += 1
		m := len(wordIndices)
		if m == maxLen {
			word := make([]byte, m)
			for i := range word {
				word[i] = alphaBytes[wordIndices[i]]
			}
			res = append(res, string(word))
		}

		for n < maxLen {
			wordIndices = append(wordIndices, wordIndices[n-m])
			n++
		}
		for len(wordIndices) > 0 && wordIndices[n-1] == alphaLen-1 {
			wordIndices = wordIndices[:n-1]
			n--
		}
	}
	return res
}
