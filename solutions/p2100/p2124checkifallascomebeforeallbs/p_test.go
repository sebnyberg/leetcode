package p2124checkifallascomebeforeallbs

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkString(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"aaabbb", true},
		{"abab", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, checkString(tc.s))
		})
	}
}

func checkString(s string) bool {
	bIdx := strings.IndexRune(s, 'b')
	return bIdx == -1 || !strings.ContainsRune(s[bIdx:], 'a')
}
