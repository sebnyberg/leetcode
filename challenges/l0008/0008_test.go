package l0008_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_l0008(t *testing.T) {
	tcs := []struct {
		in   string
		want int
	}{
		{"42", 42},
		{"", 0},
		{" ", 0},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -1 << 31},
		{"+1", 1},
		{"9223372036854775808", 2147483647},
		{"-9223372036854775808", -2147483648},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, l0008(tc.in))
		})
	}
}

func l0008(s string) int {
	ss := []rune(s)
	if len(ss) == 0 {
		return 0
	}

	var res int
	var i int
	// read whitespace
	for i < len(ss)-1 {
		switch ss[i] {
		case ' ':
			i++
			continue
		case '-', '+', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			goto EndLoop
		default:
			return res
		}
	}
EndLoop:

	// Parse sign if it exists
	var negative bool
	if ss[i] == '-' {
		negative = true
		i++
	} else if ss[i] == '+' {
		i++
	}

	// read number
	for i < len(ss) && res < 1<<31-1 && res > -1<<31 {
		n := int(ss[i] - '0')
		if n < 0 || n > 9 {
			break
		}
		res = res*10 + n
		i++
	}

	if negative {
		res = -res
	}

	if res >= 1<<31 {
		return (1 << 31) - 1
	}

	if res < -1<<31 {
		return -1 << 31
	}

	return res
}
