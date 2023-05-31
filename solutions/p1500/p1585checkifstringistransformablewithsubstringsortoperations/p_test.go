package p1585checkifstringistransformablewithsubstringsortoperations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isTransformable(t *testing.T) {
	for i, tc := range []struct {
		s    string
		t    string
		want bool
	}{
		{"84532", "34852", true},
		{"34521", "23415", true},
		{"12345", "12435", false},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, isTransformable(tc.s, tc.t))
		})
	}
}

func isTransformable(s string, t string) bool {
	// We know certain things:
	// If a digit is encountered in t that is smaller than the leftmost digit in
	// s, then there is no solution.
	// If a digit is encountered in t that is equal to the leftmost digit in s,
	// then it can be shuffled to the start of the list of values
	var idx [10][]int
	var pos [10]int
	for i := range s {
		idx[s[i]-'0'] = append(idx[s[i]-'0'], i)
	}
	for _, ch := range t {
		ch -= '0'
		if pos[ch] >= len(idx[ch]) {
			return false
		}
		for i := 0; i < int(ch); i++ {
			if pos[i] < len(idx[i]) && idx[i][pos[i]] < idx[ch][pos[ch]] {
				return false
			}
		}
		pos[ch]++
	}
	return true
}
