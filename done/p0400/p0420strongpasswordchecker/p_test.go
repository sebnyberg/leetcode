package p0420strongpasswordchecker

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_strongPasswordChecker(t *testing.T) {
	for _, tc := range []struct {
		password string
		want     int
	}{
		{"aaaaaaaAAAAAA6666bbbbaaaaaaABBC", 13},
		{"aaaaAAAAAA000000123456", 5},
		{"bbaaaaaaaaaaaaaaacccccc", 8},
		{"ABABABABABABABABABAB1", 2},
		{"1337C0d3", 0},
		{"aA1", 3},
		{"a", 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.password), func(t *testing.T) {
			require.Equal(t, tc.want, strongPasswordChecker(tc.password))
		})
	}
}

func strongPasswordChecker(password string) int {
	// The hard part is to pick inserts / replacements in such a way
	// that it minimizes the number of actions needed to reduce repeated
	// counts in a group to <= 2
	//
	// Given a count > 2, the overflow can be fixed by the actions below:
	//
	//   Count  | Replace |  Insert | Remove
	//     3    |    1    |    1    |   1
	//     4    |    1    |    1    |   2
	//     5    |    1    |    2    |   3
	//     6    |    2    |    2    |   4
	//     7    |    2    |    3    |   5
	//     8    |    2    |    4    |   6
	//
	// Conclusions:
	// * Replacements are more valuable than inserts.
	// * Value of insert/removal depends on whether it unlocks efficient replacement
	// * There is no need to mix insert and remove, use insert/remove to match
	//		length requirements, then use replace.
	//
	// Cases:
	// 1. When len(password) < 6, insert to unlock effective replacements.
	// 2. When len(password) > 20, remove to unlock effective replacements.
	// 3. When 6 <= len(password) <= 20, replacements are always best.
	//
	// An effective insert and remove reduces count to an even multiple of 3
	//
	_, counts := rle(password)
	n := len(password)

	// Match length criteria
	var nremove, ninsert int
	switch {
	case n < 6:
		for n < 6 && hasRepeats(counts) {
			// Effective inserts
			for i := 0; n < 6 && i < len(counts); i++ {
				if counts[i] <= 2 || counts[i]%3 != 1 {
					continue
				}
				n++
				ninsert++
				counts[i] -= 2
			}
			// Regular inserts
			for i := 0; n < 6 && i < len(counts); i++ {
				if counts[i] <= 2 {
					continue
				}
				n++
				ninsert++
				counts[i] -= 2
			}
		}
		// Insert remainder
		if n < 6 {
			ninsert += 6 - n
		}
	case n > 20:
		for n > 20 && hasRepeats(counts) {
			// Effective removal
			for i := 0; n > 20 && i < len(counts); i++ {
				if counts[i] <= 2 || counts[i]%3 != 0 {
					continue
				}
				n--
				nremove++
				counts[i]--
			}
			// Half-effective removal
			for i := 0; n > 21 && i < len(counts); i++ {
				if counts[i] <= 2 || counts[i]%3 != 1 {
					continue
				}
				n -= 2
				nremove += 2
				counts[i] -= 2
			}
			// Regular removal - simply remove counts until n == 20
			for i := 0; n > 20 && i < len(counts); i++ {
				if counts[i] <= 2 {
					continue
				}
				for n > 20 && counts[i] > 2 {
					n--
					nremove++
					counts[i]--
				}
			}
		}
		// Remove remainder
		if n > 20 {
			nremove += n - 20
		}
	}

	// Make replacements
	var nreplace int
	for _, c := range counts {
		nreplace += c / 3
	}

	// Count total number of actions
	constrs := countMissingConstraints(password)
	actions := nremove + max(constrs, nreplace+ninsert)
	return actions
}

func hasRepeats(counts []int) bool {
	for _, c := range counts {
		if c > 2 {
			return true
		}
	}
	return false
}

func rle(s string) ([]rune, []int) {
	chars := make([]rune, 0)
	counts := make([]int, 0)
	var n int
	for _, r := range s {
		if n > 0 && r == chars[n-1] {
			counts[n-1]++
			continue
		}
		chars = append(chars, r)
		counts = append(counts, 1)
		n++
	}
	return chars, counts
}

func countMissingConstraints(s string) int {
	var constrMissing int
	if !strings.ContainsAny(s, "0123456789") {
		constrMissing++
	}
	if !strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyz") {
		constrMissing++
	}
	if !strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		constrMissing++
	}
	return constrMissing
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
