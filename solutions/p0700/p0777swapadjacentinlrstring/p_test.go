package p0777swapadjacentinlrstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canTransform(t *testing.T) {
	for _, tc := range []struct {
		start string
		end   string
		want  bool
	}{
		{"RXXLRXRXL", "XRLXXRRLX", true},
		{"X", "L", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.start), func(t *testing.T) {
			require.Equal(t, tc.want, canTransform(tc.start, tc.end))
		})
	}
}

func canTransform(start string, end string) bool {
	// First off, letter counts must match
	var startCount, endCount [26]int
	for i := range start {
		startCount[start[i]-'A']++
		endCount[end[i]-'A']++
	}
	if startCount != endCount {
		return false
	}

	var i int
	n := len(start)
	ss := []byte(start)
	for i < n {
		if ss[i] == end[i] {
			i++
			continue
		}
		switch end[i] {
		case 'X':
			// Find an X to swap with
			j := i
			for ss[j] != 'X' {
				if ss[j] == 'L' {
					return false
				}
				j++
			}
			ss[i], ss[j] = ss[j], ss[i]
		case 'R':
			// We're screwed.
			// If it is an X, then the X can be moved forward with a superseding L,
			// but then we get stuck. The X can be swapped backwards with superseding
			// Rs, but that doesn't help either
			return false
		case 'L':
			// Find an L to swap with
			// Any R along the way means that we're screwed
			j := i
			for ss[j] != 'L' {
				if ss[j] == 'R' {
					return false
				}
				j++
			}
			ss[i], ss[j] = ss[j], ss[i]
		}
		i++
	}
	return ss[n-1] == end[n-1]
}
