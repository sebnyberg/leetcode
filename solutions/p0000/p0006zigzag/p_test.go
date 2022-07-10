package p0006zigzag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_convert(t *testing.T) {
	tcs := []struct {
		s     string
		nrows int
		want  string
	}{
		{"A", 1, "A"},
		{"AB", 1, "AB"},
		{"abcde", 2, "acebd"},
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v/%v", tc.s, tc.nrows), func(t *testing.T) {
			require.Equal(t, tc.want, convert(tc.s, tc.nrows))
		})
	}
}

func convert(s string, numRows int) string {
	// The are two kinds of rows:
	// 1. First and last rows
	// 2. Middle rows
	//
	// There are max(0, numRows-2) middle rows.
	// The middle rows require twice as many characters as the first and last
	// rows.
	// This gives us cycleLength = numRows
	cycleLength := numRows + max(0, numRows-2)

	// Top and bottom rows have one character per cycle length
	// Middle rows have two characters per cycle length: one at the nth index from
	// the start, and one at the end-nth index.

	// Time: O(n)
	// Space: O(n)

	n := len(s)
	res := make([]byte, 0, n)

	// Fill top row
	for i := 0; i < n; i += cycleLength {
		res = append(res, s[i])
	}

	// Fill middle rows
	for k := 1; k <= numRows-2; k++ {
		l := k
		r := cycleLength - k
		for l < n || r < n {
			if l < n {
				res = append(res, s[l])
			}
			if r < n {
				res = append(res, s[r])
			}
			l += cycleLength
			r += cycleLength
		}
	}

	// Fill bottom row
	if numRows > 1 {
		for i := numRows - 1; i < n; i += cycleLength {
			res = append(res, s[i])
		}
	}

	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
