package p2380timeneededtorearrangeabinarystring

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_secondsToRemoveOccurrences(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"0110101", 4},
		{"11100", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, secondsToRemoveOccurrences(tc.s))
		})
	}
}

func secondsToRemoveOccurrences(s string) int {
	curr := []byte(s)
	curr = bytes.TrimRight(curr, "0")
	curr = bytes.TrimLeft(curr, "1")
	next := make([]byte, 0, len(s))
	var steps int
	for ; len(curr) > 0; steps++ {
		next = next[:0]
		for i := 0; i < len(curr); {
			if i < len(curr)-1 && curr[i] == '0' && curr[i+1] == '1' {
				next = append(next, '1', '0')
				i += 2
			} else {
				next = append(next, curr[i])
				i++
			}
		}
		// Trim right edge
		next = bytes.TrimRight(next, "0")
		next = bytes.TrimLeft(next, "1")
		curr, next = next, curr
	}
	return steps
}
