package p1781sumofbeautyofallsubstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_beautySum(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		// {"abbccc", 5},
		{"aabcb", 5},
		// {"aabcbaa", 17},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, beautySum(tc.s))
		})
	}
}

func beautySum(s string) int {
	// There are 2^500 different substrings (too big to brute-force)
	// 1. Convert the string into a count of characters
	var freq [26]int
	maxFreq := 0
	for _, ch := range s {
		freq[ch-'a']++
		maxFreq = max(maxFreq, freq[ch-'a'])
	}
	countPerFreq := make([]int, maxFreq+1)
	for _, count := range freq {
		if count > 0 {
			countPerFreq[count]++
		}
	}

	// Assuming counts
	//  0 1 2 3 4 5
	// [0 1 0 0 2 3]
	//
	// This could be e.g. the string "abbccc"
	//
	// What is the total beauty w.r.t "ccc"?
	// We can start by erasing 'a' and get "bbccc" -> 3-2 = 1
	// Then "bccc" -> 3-1 = 2
	// Then "abccc" -> 3-1 = 2
	// Then "accc" -> 3-1 = 2

	// So, given a maximum count number, the total beauty is going to be
	// the number of permutations for a count which has at least one character.
	// This gives us (1 << (n_per_group) - 1)

	// So, given "abbccc", the beauty for substrings containing "ccc" is
	// 3-2 * ((1 << countPerFreq[2])-1)
	// + (3-1)*((1 << (countPerFreq[2] + countPerFreq[1]))-1)
	// = 1 * (2-1) + 2 * (4-1) = 1 + 3
	// This seems right:
	// "abbccc", "abccc", "bccc", "accc"

	// What if there were multiple counts on a given frequency?
	// Then the other counts from the current level could be used at other levels

	// Let's try this brute-force:
	// "abbcc"
	// "abcc", "acc", "bcc",
	// "acbb", "abb", "cbb",

	var totalBeauty int
	for i := maxFreq; i >= 2; i-- {
		cumSum := countPerFreq[i] - 1
		beautyForLevel := 0
		for j := i - 1; j >= 1; j-- {
			cumSum += countPerFreq[j]
			levelPerms := ((1 << cumSum) - 1)
			beautyForCount := (i - j) * levelPerms
			beautyForLevel += beautyForCount
		}
		totalBeauty += countPerFreq[i] * beautyForLevel
		countPerFreq[i-1] += countPerFreq[i]
	}
	return totalBeauty
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
