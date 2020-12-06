package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_l0006(t *testing.T) {
	tcs := []struct {
		s     string
		nrows int
		want  string
	}{
		{"A", 1, "A"},
		{"abcde", 2, "acebd"},
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v/%v", tc.s, tc.nrows), func(t *testing.T) {
			require.Equal(t, tc.want, l0006convert(tc.s, tc.nrows))
		})
	}
}

func l0006convert(s string, nrows int) string {
	if nrows == 1 {
		return s
	}

	ss := []rune(s)

	//         (firstcol) (diagonal)
	nperzigzag := nrows + nrows - 2
	nzigzags := len(ss) / nperzigzag

	// Create a two-dimensional array that will contain the letters
	zigzagRows := make([][]rune, nrows)
	for i := range zigzagRows {
		// If this is the first or last row, the number of characters per
		// row is at most nzigzag + 1
		if i == 0 || i == nrows-1 {
			zigzagRows[i] = make([]rune, 0, nzigzags+1)
			continue
		}
		// For rows with diagonals, each row is at most 2*nzigzag+1
		zigzagRows[i] = make([]rune, 0, 2*nzigzags+1)
	}

	var i int
	for i < len(s) {
		// Iterate across first column
		for j := 0; j < nrows-1 && i < len(s); j++ {
			zigzagRows[j] = append(zigzagRows[j], ss[i])
			i++
		}

		// Iterate up along diagonal
		for j := nrows - 1; j > 0 && i < len(s); j-- {
			zigzagRows[j] = append(zigzagRows[j], ss[i])
			i++
		}
	}

	result := make([]rune, 0, len(ss))
	for _, row := range zigzagRows {
		result = append(result, row...)
	}

	return string(result)
}
