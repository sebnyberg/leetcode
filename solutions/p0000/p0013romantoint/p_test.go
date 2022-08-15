package p0013romantoint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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

var deltaOne [22]int16 = [22]int16{
	'I' - 'C': 1,
	'V' - 'C': 5,
	'X' - 'C': 10,
	'L' - 'C': 50,
	'C' - 'C': 100,
	'D' - 'C': 500,
	'M' - 'C': 1000,
}

var deltaTwo [22][22]int16 = [22][22]int16{
	'I' - 'C': {'X' - 'C': -2, 'V' - 'C': -2},
	'X' - 'C': {'L' - 'C': -20, 'C' - 'C': -20},
	'C' - 'C': {'D' - 'C': -200, 'M' - 'C': -200},
}

func romanToInt(s string) int {
	var res int16
	for i := range s {
		res += deltaOne[s[i]-'C']
	}
	for i := 1; i < len(s); i++ {
		res += deltaTwo[s[i-1]-'C'][s[i]-'C']
	}
	return int(res)
}
