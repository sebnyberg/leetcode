package l0014_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestCommonPrefix(t *testing.T) {
	for _, tc := range []struct {
		in   []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{}, ""},
		{[]string{"dog", "racecar", "car"}, ""},
	} {
		t.Run(fmt.Sprintf("+%v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, longestCommonPrefix(tc.in))
		})
	}
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	prefix := strs[0]
	var j int
	for i := 1; i < len(strs); i++ {
		prefixlen := min(len(prefix), len(strs[i]))
		for j = 0; j < prefixlen && strs[i][j] == prefix[j]; j++ {
		}
		prefix = prefix[:j]
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}
