package p0423reconstructdigitsfromeng

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_originalDigits(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"owoztneoer", "012"},
		{"fviefuro", "45"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, originalDigits(tc.s))
		})
	}
}

func originalDigits(s string) string {
	// zero: 'z'
	// two: 'w'
	// four: 'u'
	// one: 'o'
	// six: 'x'
	// eight: 'g'
	// three: 't' after removing 'two'
	// five: 'f' after four
	// nines: 'i' after all others
	// seven: 's' after all others
	type matcher struct {
		ch     rune
		numChs []rune
		num    byte
	}
	matchers := []matcher{
		{'z', []rune{'z', 'e', 'r', 'o'}, '0'},
		{'w', []rune{'t', 'w', 'o'}, '2'},
		{'u', []rune{'f', 'o', 'u', 'r'}, '4'},
		{'x', []rune{'s', 'i', 'x'}, '6'},
		{'g', []rune{'e', 'i', 'g', 'h', 't'}, '8'},
		{'o', []rune{'o', 'n', 'e'}, '1'},
		{'t', []rune{'t', 'h', 'r', 'e', 'e'}, '3'},
		{'f', []rune{'f', 'i', 'v', 'e'}, '5'},
		{'i', []rune{'n', 'i', 'n', 'e'}, '9'},
		{'s', []rune{'s', 'e', 'v', 'e', 'n'}, '7'},
	}

	res := make([]byte, 0, len(s)/3)
	var chCount [26]int
	for _, ch := range s {
		chCount[ch-'a']++
	}
	for _, m := range matchers {
		for chCount[m.ch-'a'] > 0 {
			for _, r := range m.numChs {
				chCount[r-'a']--
			}
			res = append(res, m.num)
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return string(res)
}
