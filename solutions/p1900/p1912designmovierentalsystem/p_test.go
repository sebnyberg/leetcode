package p1912designmovierentalsystem

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	newlineSep = []byte{'\n'}
	commaSep   = []byte{','}
)

func TestNothing(t *testing.T) {
	for _, tc := range []struct {
		infile  string
		outfile string
	}{
		{"testdata/1in", "testdata/1out"},
		{"testdata/2in", "testdata/2out"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.infile), func(t *testing.T) {
			inF, _ := os.Open(tc.infile)
			inB, _ := io.ReadAll(inF)
			parts := bytes.Split(inB, newlineSep)
			actions := parseRootParts(parts[0], commaSep)
			inputs := parseRootParts(parts[1], commaSep)
			outF, _ := os.Open(tc.outfile)
			outB, _ := io.ReadAll(outF)
			firstOutputRow := bytes.Split(outB, newlineSep)[0]
			outputs := parseRootParts(firstOutputRow, commaSep)
			c := construct(inputs[0])
			n := len(inputs)
			for i := 1; i < n; i++ {
				inputParts := parseRootParts(inputs[i], commaSep)
				switch string(actions[i]) {
				case `"search"`:
					movie, _ := strconv.Atoi(string(inputParts[0]))
					res := c.Search(movie)
					want := make([]int, 0, 5)
					for _, shopBytes := range parseRootParts(outputs[i], commaSep) {
						x, _ := strconv.Atoi(string(shopBytes))
						want = append(want, x)
					}
					require.Equal(t, want, res, i)
				case `"rent"`:
					shop, _ := strconv.Atoi(string(inputParts[0]))
					movie, _ := strconv.Atoi(string(inputParts[1]))
					c.Rent(shop, movie)
				case `"drop"`:
					shop, _ := strconv.Atoi(string(inputParts[0]))
					movie, _ := strconv.Atoi(string(inputParts[1]))
					c.Drop(shop, movie)
				case `"report"`:
					res := c.Report()
					want := make([][]int, 0)
					for _, entry := range parseRootParts(outputs[i], commaSep) {
						res := make([]int, 0)
						for _, entryPart := range parseRootParts(entry, commaSep) {
							x, _ := strconv.Atoi(string(entryPart))
							res = append(res, x)
						}
						want = append(want, res)
					}
					require.Equal(t, want, res, i)
				}
			}
			_ = outputs
		})
	}
}

func construct(bs []byte) MovieRentingSystem {
	parts := parseRootParts(bs, commaSep)
	n, _ := strconv.Atoi(string(parts[0]))
	entryBytes := parseRootParts(parts[1], commaSep)
	entries := make([][]int, len(entryBytes))
	for i, eb := range entryBytes {
		entries[i] = make([]int, 3)
		for j, ebb := range parseRootParts(eb, commaSep) {
			nn, _ := strconv.Atoi(string(ebb))
			entries[i][j] = nn
		}
	}
	return Constructor(n, entries)
}

func parseRootParts(s []byte, delim []byte) [][]byte {
	s = s[1 : len(s)-1]
	// Split on all root-level commas
	parts := make([][]byte, 0)
	stack := []byte{}
	var pos int
	for i, b := range s {
		if b == '[' {
			stack = append(stack, b)
			continue
		}
		if b == ']' {
			stack = stack[:len(stack)-1]
			continue
		}
		if b == ',' && len(stack) == 0 {
			parts = append(parts, s[pos:i])
			pos = i + 1
		}
	}
	if len(s)-pos >= 1 {
		parts = append(parts, s[pos:])
	}
	return parts
}
