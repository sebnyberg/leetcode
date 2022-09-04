package p2399checkdistancesbetweensameletters

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkDistances(t *testing.T) {
	for _, tc := range []struct {
		s        string
		distance []int
		want     bool
	}{
		{"abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, checkDistances(tc.s, tc.distance))
		})
	}
}

func checkDistances(s string, distance []int) bool {
	// Store the position of the first occurrence of each character and verify
	// that the distance to the second occurrence is correct.
	var pos [26]int
	for x := range pos {
		pos[x] = -1
	}
	for i, ch := range s {
		c := int(ch - 'a')
		if pos[c] == -1 {
			pos[c] = i
		} else {
			if i-pos[c]-1 != distance[c] {
				return false
			}
		}
	}
	return true
}
