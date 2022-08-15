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

var deltaOne [26]int16 = [26]int16{
	'I' - 'A': 1,
	'V' - 'A': 5,
	'X' - 'A': 10,
	'L' - 'A': 50,
	'C' - 'A': 100,
	'D' - 'A': 500,
	'M' - 'A': 1000,
}

var deltaTwo [26][26]int16 = [26][26]int16{
	'I' - 'A': {'X' - 'A': -2, 'V' - 'A': -2},
	'X' - 'A': {'L' - 'A': -20, 'C' - 'A': -20},
	'C' - 'A': {'D' - 'A': -200, 'M' - 'A': -200},
}

func romanToInt(s string) int {
	var res int16
	for i := range s {
		res += deltaOne[s[i]-'A']
	}
	for i := 1; i < len(s); i++ {
		res += deltaTwo[s[i-1]-'A'][s[i]-'A']
	}
	return int(res)
}
