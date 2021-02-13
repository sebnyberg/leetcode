package p0017letterphonenum

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_letterCombinations(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want []string
	}{
		{"234", []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}},
		// {"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		// {"", []string{}},
		// {"2", []string{"a", "b", "c"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			sort.Strings(tc.want)
			res := letterCombinations(tc.in)
			sort.Strings(res)
			require.Equal(t, tc.want, res)
		})
	}
}

var letters = map[rune][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

// Return all possible letter combinations that the number could represent
// in any order.
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	nrepeat := make([]int, len(digits))
	total := 1
	for i, digit := range digits {
		nrepeat[i] = total
		total *= len(letters[digit])
	}

	res := make([]string, total)
	for i := 0; i < total; i++ {
		for j, digit := range digits {
			res[i] = res[i] + letters[digit][i/nrepeat[j]%len(letters[digit])]
		}
	}
	return res
}
