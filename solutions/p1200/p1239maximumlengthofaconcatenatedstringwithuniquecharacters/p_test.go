package p2439maximumlengthofaconcatenatedstringwithuniquecharacters

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxLength(t *testing.T) {
	for i, tc := range []struct {
		arr  []string
		want int
	}{
		{[]string{"un", "iq", "ue"}, 4},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxLength(tc.arr))
		})
	}
}

func maxLength(arr []string) int {
	// Note that while we may pick any number of strings, we cannot reorder them
	//
	// That is, a pair of strings does not result in two concatenations.
	//
	// One approach that I like is to figure out some element-specific
	// invariant. "Assuming that we visit each element from left to
	// right, what would we need to record on each visit?".
	//
	// The current string must be added to the list of options on its own,
	// and in combination with any prior combination of strings that does not
	// overlap with the current string.
	m := make(map[int]struct{})
	m[0] = struct{}{}
	for _, s := range arr {
		var b int
		for _, ch := range s {
			p := int(1 << (ch - 'a'))
			if p&b > 0 {
				goto next
			}
			b |= p
		}
		for k := range m {
			if k&b == 0 {
				m[k|b] = struct{}{}
			}
		}
	next:
	}
	var res int
	for k := range m {
		res = max(res, bits.OnesCount(uint(k)))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
