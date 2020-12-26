package d26_test

import (
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func Test_numDecodings(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want int
	}{
		{"12", 2},
		{"226", 3},
		{"0", 0},
		{"1", 1},
		{"2101", 1},
	} {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, numDecodings(tc.in))
		})
	}
}

type memoizedDecoder struct {
	decodeCount map[string]int
}

func numDecodings(s string) int {
	var start int

	// Skip leading zeroes
	ch, _ := utf8.DecodeRuneInString(s)
	if ch == '0' {
		return 0
	}

	dec := memoizedDecoder{
		decodeCount: make(map[string]int, len(s[start:])),
	}

	dec.decodeCount["0"] = 0
	dec.decodeCount[""] = 1

	return dec.numDecodings(s)
}

func (d memoizedDecoder) numDecodings(s string) int {
	if decodeCount, exists := d.decodeCount[s]; exists {
		return decodeCount
	}

	switch len(s) {
	case 0:
		d.decodeCount[s] = 0
		return 1
	case 1:
		d.decodeCount[s] = 1
		return 1
	}

	ch, width := utf8.DecodeRuneInString(s)
	if ch == '0' {
		return 0
	}

	// ch is 1-9, calculate decoding results
	// from shifting one digit
	res := d.numDecodings(s[width:])
	d.decodeCount[s[width:]] = res

	// read second character
	secondCh, secondWidth := utf8.DecodeRuneInString(s[width:])

	switch {
	case ch > '2' || ch == '2' && secondCh > '6':
		return res
	case ch == '1' || ch == '2': // secondCh is implicitly <= 6
		res2 := d.numDecodings(s[width+secondWidth:])
		d.decodeCount[s[width+secondWidth:]] = res2
		return res + res2
	default:
		panic("invalid ch")
	}
}
