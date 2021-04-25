package p1784checkifbinarystringhasatmostonesegmentofones

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkOnesSegment(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"1001", false},
		{"110", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, checkOnesSegment(tc.s))
		})
	}
}

func checkOnesSegment(s string) bool {
	first := strings.IndexRune(s, '1')
	if first == -1 {
		return true
	}
	zeroAfterFirst := strings.IndexRune(s[first:], '0')
	if zeroAfterFirst == -1 {
		return true
	}
	second := strings.IndexRune(s[first+zeroAfterFirst:], '1')
	return second == -1
}
