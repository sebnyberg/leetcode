package p0336palindromepairs

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func Test_palindromePairs(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  [][]int
	}{
		{[]string{"abcd", "dcba", "lls", "s", "sssll"}, [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}}},
		{[]string{"bat", "tab", "cat"}, [][]int{{0, 1}, {1, 0}}},
		{[]string{"a", ""}, [][]int{{0, 1}, {1, 0}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, palindromePairs(tc.words))
		})
	}
}

var res [][]int

func BenchmarkPalindromePairs(b *testing.B) {
	f, _ := os.Open("testdata/input")
	bs, _ := io.ReadAll(f)
	words := make([]string, 0)
	for _, quotedWord := range strings.Split(string(bs), ",") {
		words = append(words, quotedWord[1:len(quotedWord)-1])
	}
	for i := 0; i < b.N; i++ {
		b := palindromePairs(words)
		res = b
	}
}

func palindromePairs(words []string) [][]int {
	n := len(words)
	wordIndices := make(map[string][]uint16, n)
	// starWordIndices are words for which the first letter can be anything
	midPalinWordIndices := make(map[string][]uint16, n)
	revWords := make([][]byte, n)
	for wordIdx, word := range words {
		wordIndices[word] = append(wordIndices[word], uint16(wordIdx))
		bs := []byte(word)
		m := len(word)
		// For each position in the word, try to make a palindrome and add
		// the remainder to the midPalinWordIndices.
		for i := 1; i <= m; i++ {
			if isPalin(bs[:i]) {
				midPalinWordIndices[word[i:]] = append(midPalinWordIndices[word[i:]], uint16(wordIdx))
			}
		}
		// Store reverse in revWords
		for l, r := 0, m-1; l < r; l, r = l+1, r-1 {
			bs[l], bs[r] = bs[r], bs[l]
		}
		revWords[wordIdx] = bs
	}

	res := make([][]int, 0)
	rhs := make(map[uint16]struct{})
	for i := range words {
		m := len(words[i])
		rev := revWords[i]
		// right side has palindrome in the middle
		for _, idx := range midPalinWordIndices[toStr(rev)] {
			rhs[idx] = struct{}{}
		}
		// this side has palindrome in the middle
		for i := 0; i <= m; i++ {
			if isPalin(rev[:i]) {
				for _, idx := range wordIndices[toStr(rev[i:])] {
					rhs[idx] = struct{}{}
				}
			}
		}
		delete(rhs, uint16(i)) // in case there was a self-reference
		for rhsIdx := range rhs {
			res = append(res, []int{i, int(rhsIdx)})
		}
		for k := range rhs {
			delete(rhs, k)
		}
	}
	return res
}

func toStr(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func isPalin(bs []byte) bool {
	for l, r := 0, len(bs)-1; l < r; l, r = l+1, r-1 {
		if bs[l] != bs[r] {
			return false
		}
	}
	return true
}
