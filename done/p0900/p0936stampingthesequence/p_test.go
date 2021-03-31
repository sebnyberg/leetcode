package p0936stampingthesequence

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_movesToStamp(t *testing.T) {
	for _, tc := range []struct {
		stamp  string
		target string
		want   []int
	}{
		{"zbs", "zbzbsbszbssbzbszbsss", []int{5, 2, 8, 6, 3, 0, 7, 4, 1}},
		{"de", "ddeddeddee", []int{5, 2, 8, 6, 3, 0, 7, 4, 1}},
		{"abc", "ababc", []int{0, 2}},
		// {"abca", "aabcaca", []int{3, 0, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.stamp), func(t *testing.T) {
			require.EqualValues(t, tc.want, movesToStamp(tc.stamp, tc.target))
		})
	}
}

func movesToStamp(stamp string, target string) []int {
	indicesToExplore := make([]int, 0)
	stampedIndices := make([]int, 0)
	targetBytes := []byte(target)

	place := func(indexToStamp int) {
		stampedIndices = append(stampedIndices, indexToStamp)
		// Place stamp
		for i := 0; i < len(stamp) && indexToStamp+i < len(target); i++ {
			if indexToStamp < 0 {
				continue
			}
			targetBytes[indexToStamp+i] = 0
		}
	}

	idx := strings.Index(target, stamp)
	for offset := 0; idx != -1; {
		indicesToExplore = append(indicesToExplore, offset+idx)
		place(offset + idx)
		offset += idx + len(stamp)
		idx = strings.Index(target[offset:], stamp)
	}
	for len(indicesToExplore) > 0 {
		newIndicesToExplore := make([]int, 0)
		for _, indexToExplore := range indicesToExplore {
			// Find right stamp
			for l := indexToExplore + len(stamp); l > indexToExplore; l-- {
				if match(targetBytes[l:], stamp) {
					place(l)
					newIndicesToExplore = append(newIndicesToExplore, l)
					break
				}
			}

			// Find left stamp
			for l := indexToExplore - len(stamp); l < indexToExplore; l++ {
				if l < 0 {
					continue
				}
				if match(targetBytes[l:], stamp) {
					place(l)
					newIndicesToExplore = append(newIndicesToExplore, l)
					break
				}
			}
			indicesToExplore = newIndicesToExplore
		}
	}

	for _, b := range targetBytes {
		if b != 0 {
			return []int{}
		}
	}

	// Reverse stamps
	for l, r := 0, len(stampedIndices)-1; l < r; l, r = l+1, r-1 {
		stampedIndices[l], stampedIndices[r] = stampedIndices[r], stampedIndices[l]
	}

	return stampedIndices
}

// Match the provided pattern against the provided bytes
// Zeroes in b are considered to match anything in the pattern
// Returns false if len(pat) > len(b)
func match(b []byte, pat string) bool {
	if len(b) < len(pat) {
		return false
	}
	allZeroes := true
	for i := range pat {
		if b[i] == 0 {
			continue
		}
		if b[i] != pat[i] {
			return false
		}
		allZeroes = false
	}
	return !allZeroes
}
