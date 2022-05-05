package p0844backspacestringcompare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_backspaceCompare(t *testing.T) {
	for _, tc := range []struct {
		S    string
		T    string
		want bool
	}{
		// {"ab#c", "ad#c", true},
		// {"ab##", "c#d#", true},
		// {"a##c", "#a#c", true},
		{"a#c", "b", false},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.S, tc.T), func(t *testing.T) {
			require.Equal(t, tc.want, backspaceCompare(tc.S, tc.T))
		})
	}
}

func backspaceCompare(S string, T string) bool {
	sstack := make([]byte, 0)
	for i := range S {
		if S[i] == '#' {
			if len(sstack) > 0 {
				sstack = sstack[:len(sstack)-1]
			}
		} else {
			sstack = append(sstack, S[i])
		}
	}
	tstack := make([]byte, 0)
	for i := range T {
		if T[i] == '#' {
			if len(tstack) > 0 {
				tstack = tstack[:len(tstack)-1]
			}
		} else {
			tstack = append(tstack, T[i])
		}
	}
	if len(sstack) != len(tstack) {
		return false
	}
	for i := range sstack {
		if sstack[i] != tstack[i] {
			return false
		}
	}
	return true
}
