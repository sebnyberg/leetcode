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

func Test_createIndices(t *testing.T) {
	tcs := []struct {
		nrows int
		ntot  int
		want  []int
	}{
		{1, 2, []int{0, 1}},
		{2, 2, []int{0, 1}},
		{1, 5, []int{0, 1, 2, 3, 4}},
		{2, 5, []int{0, 2, 4, 1, 3}},
		{3, 5, []int{0, 4, 1, 3, 2}},
		{3, 14, []int{0, 4, 8, 12, 1, 3, 5, 7, 9, 11, 13, 2, 6, 10}},
		{4, 5, []int{0, 1, 2, 4, 3}},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v/%v", tc.nrows, tc.ntot), func(t *testing.T) {
			require.Equal(t, tc.want, createIndices(tc.nrows, tc.ntot))
		})
	}
}

func createIndices(nrows int, ntot int) []int {
	res := make([]int, ntot)
	if nrows == 1 {
		for i := range res {
			res[i] = i
		}
		return res
	}
	nperzigzag := nrows + nrows - 2

	// first row
	var i int
	for j := 0; j < ntot; j += nperzigzag {
		res[i] = j
		i++
	}

	// intermediate rows
	for row := 1; row < nrows-1; row++ {
		// current zigzag
		j := 0
		for {
			if row+j*nperzigzag >= ntot {
				break
			}
			res[i] = row + j*nperzigzag
			i++
			j++
			if j*nperzigzag-row >= ntot {
				break
			}
			res[i] = j*nperzigzag - row
			i++
		}
	}

	// last row
	for j := nrows - 1; j < ntot; j += nperzigzag {
		res[i] = j
		i++
	}

	return res
}

func convert(s string, nrows int) string {
	idx := createIndices(nrows, len(s))
	res := make([]rune, len(s))
	for i, idx := range idx {
		res[i] = rune(s[idx])
	}
	return string(res)
}
