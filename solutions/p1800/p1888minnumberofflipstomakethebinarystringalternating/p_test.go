package p1888minnumberofflipstomakethebinarystringalternating

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minFlips(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"01001001101", 2},
		{"111000", 2},
		{"010", 0},
		{"1110", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, minFlips(tc.s))
		})
	}
}

func minFlips(s string) int {
	zeroes := []int{0, 0}
	ones := []int{0, 0}
	// make initial count of zeroes at even and odd positions
	for i, ch := range s {
		if ch == '0' {
			zeroes[i%2]++
		} else {
			ones[i%2]++
		}
	}

	minResult := min(ones[0]+zeroes[1], zeroes[0]+ones[1])
	if len(s)%2 == 0 {
		return minResult
	}

	// the number of values is odd, we can shift to alter the makeup of the list
	for _, ch := range s {
		if ch == '0' {
			// Move the zero from even to odd (shift the number to the end)
			zeroes[1]++
			zeroes[0]--
		} else {
			ones[1]++
			ones[0]--
		}
		zeroes[0], zeroes[1] = zeroes[1], zeroes[0]
		ones[0], ones[1] = ones[1], ones[0]
		minResult = min(minResult, min(ones[0]+zeroes[1], zeroes[0]+ones[1]))
	}
	return minResult
}

// Count flips using strategy #2 only
func countFlips(bs []byte, zeroFirst bool) int {
	var flips int
	zero := 0
	if !zeroFirst {
		zero = 1
	}
	for i := 0; i < len(bs); i++ {
		if i%2 == zero {
			if bs[i] != '0' {
				flips++
			}
		} else {
			if bs[i] != '1' {
				flips++
			}
		}
	}
	return flips
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
