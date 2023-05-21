package p1307verbalarithmeticpuzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSolvable(t *testing.T) {
	for i, tc := range []struct {
		words  []string
		result string
		want   bool
	}{
		{[]string{"WE", "ARE"}, "IT", false},
		{[]string{"A", "B"}, "A", true},
		{[]string{"CBA", "CBA", "CBA", "CBA", "CBA"}, "EDD", false},
		{[]string{"SEND", "MORE"}, "MONEY", true},
		{[]string{"SIX", "SEVEN", "SEVEN"}, "TWENTY", true},
		{[]string{"LEET", "CODE"}, "POINT", false},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, isSolvable(tc.words, tc.result))
		})
	}
}

func isSolvable(words []string, result string) bool {
	// We can perform DFS to check whether there exists a solution.
	// For each position in result, from right to left, try all possible
	// character mappings for the letters found in that position in the words
	//
	// If this for some reason is too slow, we may also keep track of duplicate
	// guesses.
	//

	var charMap [256]byte
	var nextChar byte = 'A'

	// normalize maps characters to [A,(A+9)] to make them 0-9 indexable, and
	// reverses strings so that addition goes from left to right instead of
	// right to left.
	normalize := func(s string) string {
		bs := []byte(s)
		for l, r := 0, len(bs)-1; l < r; l, r = l+1, r-1 {
			bs[l], bs[r] = bs[r], bs[l]
		}
		for i := range bs {
			if charMap[bs[i]] == 0 {
				charMap[bs[i]] = nextChar
				nextChar++
			}
			bs[i] = charMap[bs[i]]
		}
		return string(bs)
	}
	_ = normalize

	for i := range words {
		words[i] = normalize(words[i])
	}
	result = normalize(result)

	var m [10]int
	for i := range m {
		m[i] = -1
	}
	var assigned [10]bool
	res := dfs(words, result, m, assigned, 0, 0, 0)
	return res
}

func isLastDigit(w string, i int) bool {
	if len(w) == 1 {
		return false
	}
	for j := i + 1; j < len(w); j++ {
		if w[j] != w[i] {
			return false
		}
	}
	return true
}

func dfs(words []string, result string, m [10]int, assigned [10]bool, i, j, val int) bool {
	// m[c] = value of character c
	// assigned[x] = whether value x has been assigned to a character
	// i = character position
	// j = current word
	// val = current value

	if i >= len(result) && j >= len(words) {
		// There is an annoying test-case A + A = AA where A = 0 is invalid
		// because we can't have two zeroes in the result. Hence this weird
		// check.
		return val == 0 && m[result[len(result)-1]-'A'] == 0
	}

	if j >= len(words) {
		// Validate the current digit against the character in the result
		digit := val % 10
		if m[result[i]-'A'] == -1 {
			if assigned[digit] {
				return false
			}
			if digit == 0 && isLastDigit(result, i) {
				return false
			}
			m[result[i]-'A'] = digit
			assigned[digit] = true
		}
		if m[result[i]-'A'] != digit {
			return false
		}
		return dfs(words, result, m, assigned, i+1, 0, val/10)
	}

	w := words[j]
	if i >= len(w) {
		// Nothing to do here
		return dfs(words, result, m, assigned, i, j+1, val)
	}

	if m[w[i]-'A'] != -1 {
		// An assignment has been made already, use it
		return dfs(words, result, m, assigned, i, j+1, val+m[w[i]-'A'])
	}

	// Try all available assignments for this character.
	for x := range assigned {
		if assigned[x] {
			continue
		}
		if x == 0 && isLastDigit(w, i) {
			continue
		}
		assigned[x] = true
		m[w[i]-'A'] = x
		ok := dfs(words, result, m, assigned, i, j+1, val+x)
		if ok {
			return true
		}
		m[w[i]-'A'] = -1
		assigned[x] = false
	}
	return false
}
