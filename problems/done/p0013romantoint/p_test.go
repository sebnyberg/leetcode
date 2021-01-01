package p0013romantoint

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_romanToInt(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	} {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, romanToInt(tc.in))
		})
	}
}

const (
	romanI int = 1
	romanV int = 5
	romanX int = 10
	romanL int = 50
	romanC int = 100
	romanD int = 500
	romanM int = 1000
)

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	var res int
	prev := romanM
	for _, ch := range s {
		switch ch {
		case 'I':
			res += romanI
			prev = romanI
		case 'V':
			res += romanV
			if prev == romanI {
				res -= 2
			}
			prev = romanV
		case 'X':
			res += romanX
			if prev == romanI {
				res -= 2
			}
			prev = romanX
		case 'L':
			res += romanL
			if prev == romanX {
				res -= 20
			}
			prev = romanL
		case 'C':
			res += romanC
			if prev == romanX {
				res -= 20
			}
			prev = romanC
		case 'D':
			res += romanD
			if prev == romanC {
				res -= 200
			}
			prev = romanD
		case 'M':
			res += romanM
			if prev == romanC {
				res -= 200
			}
			prev = romanM
		}
	}
	return res
}
