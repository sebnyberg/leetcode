package p1061lexicographicallysmalleststring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestEquivalentString(t *testing.T) {
	for _, tc := range []struct {
		A    string
		B    string
		S    string
		want string
	}{
		{"parker", "morris", "parser", "makkek"},
	} {
		t.Run(fmt.Sprintf("%v/%v/%v", tc.A, tc.B, tc.S), func(t *testing.T) {
			require.Equal(t, tc.want, smallestEquivalentString(tc.A, tc.B, tc.S))
		})
	}
}

func smallestEquivalentString(A string, B string, S string) string {
	var parent [26]byte
	for i := range parent {
		parent[i] = byte(i)
	}
	find := func(ch byte) byte {
		for parent[ch] != ch {
			ch = parent[ch]
		}
		return ch
	}

	for i := range A {
		parent[find(A[i]-'a')] = find(B[i] - 'a') // Union
	}
	// Find smallest in each group
	var smallest [26]byte
	for i := range smallest {
		smallest[i] = byte(i)
	}
	for i := range parent {
		a := byte(i)
		if a < smallest[find(a)] {
			smallest[find(a)] = a
		}
	}

	res := make([]byte, len(S))
	for i := range S {
		res[i] = smallest[find(S[i]-'a')] + 'a'
	}
	s := string(res)
	return s
}
