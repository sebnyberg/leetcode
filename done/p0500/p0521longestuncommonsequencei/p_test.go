package p0521longestuncommonsequencei

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLUSlength(t *testing.T) {
	for _, tc := range []struct {
		a    string
		b    string
		want int
	}{
		{"aefawfawfawfaw", "aefawfeawfwafwaef", 17},
		{"aba", "cdc", 3},
		{"aaa", "bbb", 3},
		{"aaa", "aaa", -1},
		{"aaa", "aaaa", 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.a), func(t *testing.T) {
			require.Equal(t, tc.want, findLUSlength(tc.a, tc.b))
		})
	}
}

func findLUSlength(a string, b string) int {
	if len(a) != len(b) {
		if len(b) > len(a) {
			return len(b)
		}
		return len(a)
	}
	for i := range a {
		if a[i] != b[i] {
			return len(a) - i
		}
	}

	return -1
}
