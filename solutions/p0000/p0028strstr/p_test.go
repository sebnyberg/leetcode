package p0028strstr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_strStr(t *testing.T) {
	for _, tc := range []struct {
		haystack string
		needle   string
		want     int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.needle, tc.haystack), func(t *testing.T) {
			require.Equal(t, tc.want, strStr(tc.haystack, tc.needle))
		})
	}
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	n, m := len(haystack), len(needle)
	for i := 0; i < n-m+1; i++ {
		if haystack[i:i+m] == needle {
			return i
		}
	}

	return -1
}
