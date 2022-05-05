package p1933checkifstringisdecomposableintovalueequalsubstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isDecomposable(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"000111000", false},
		{"00011111222", true},
		{"011100022233", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, isDecomposable(tc.s))
		})
	}
}

func isDecomposable(s string) bool {
	n := len(s)
	curCount := 1
	var hasTwoGroup bool
	ok := func(val int) bool {
		r := val % 3
		if r == 2 && !hasTwoGroup {
			if hasTwoGroup {
				return false
			}
			hasTwoGroup = true
			return true
		}
		return !(r == 1)
	}
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			if !ok(curCount) {
				return false
			}
			curCount = 0
		}
		curCount++
	}
	return ok(curCount) && hasTwoGroup
}
